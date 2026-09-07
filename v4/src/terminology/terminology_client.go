package terminology

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// TerminologyClient — HTTP client mandiri untuk SmarTerm (terminology.ivanmd.id).
//
// ⚠️ TERPISAH dari ekosistem SATUSEHAT: jangan reuse OAuth2Client SATUSEHAT
// Kemkes. SmarTerm punya auth sendiri, dual-mode:
//
//  1. Sanctum Personal Access Token (PAT) — token statis, header `Bearer`.
//     Env: TERMINOLOGY_API_TOKEN (atau argumen token).
//  2. OIDC client_credentials ke auth.ivanmd.id (Passport) — client fetch JWT
//     sendiri, cache sampai expire, auto-refresh saat 401.
//     Env: TERMINOLOGY_CLIENT_ID / TERMINOLOGY_CLIENT_SECRET /
//     TERMINOLOGY_TOKEN_URL (default https://auth.ivanmd.id/oauth/token)
//
// Jika keduanya diberikan, PAT menang (lebih sederhana). OIDC dipakai saat
// token kosong dan kredensial client tersedia.
//
// Fluent chaining:
//
//	term.Query("pemasangan kateter").On("kptl").Limit(10).Get()
type TerminologyClient struct {
	baseURL      string
	token        string
	clientID     string
	clientSecret string
	tokenURL     string
	httpClient   *http.Client

	oidcToken     string
	oidcExpiresAt int64
}

// NewTerminologyClient membuat client; baseURL/token kosong dibaca dari env
// TERMINOLOGY_BASE_URL / TERMINOLOGY_API_TOKEN, default https://terminology.ivanmd.id.
// Timeout HTTP 30 detik.
func NewTerminologyClient(baseURL, token string) *TerminologyClient {
	if baseURL == "" {
		baseURL = os.Getenv("TERMINOLOGY_BASE_URL")
	}
	if baseURL == "" {
		baseURL = "https://terminology.ivanmd.id"
	}
	if token == "" {
		token = os.Getenv("TERMINOLOGY_API_TOKEN")
	}

	return &TerminologyClient{
		baseURL:      strings.TrimRight(baseURL, "/"),
		token:        token,
		clientID:     os.Getenv("TERMINOLOGY_CLIENT_ID"),
		clientSecret: os.Getenv("TERMINOLOGY_CLIENT_SECRET"),
		tokenURL:     envOr("TERMINOLOGY_TOKEN_URL", "https://auth.ivanmd.id/oauth/token"),
		httpClient:   &http.Client{Timeout: 30 * time.Second},
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// Query memulai fluent query builder.
func (c *TerminologyClient) Query(text string) *TerminologyQueryBuilder {
	return &TerminologyQueryBuilder{client: c, text: text, category: "all", limit: 20}
}

// doJSON helper HTTP internal — set header Authorization Bearer, parse body
// JSON ke map[string]any. Error jika status >= 400.
func (c *TerminologyClient) doJSON(method, path string, query url.Values, body any) (map[string]any, error) {
	u := c.baseURL + "/" + strings.TrimLeft(path, "/")
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	token, err := c.resolveToken()
	if err != nil {
		return nil, err
	}

	send := func(tok string) (map[string]any, int, error) {
		var rdr io.Reader
		if body != nil {
			b, err := json.Marshal(body)
			if err != nil {
				return nil, 0, err
			}
			rdr = strings.NewReader(string(b))
		}
		req, err := http.NewRequest(method, u, rdr)
		if err != nil {
			return nil, 0, err
		}
		req.Header.Set("Accept", "application/json")
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Header.Set("User-Agent", "satusehat-integration-go/terminology-client")
		if tok != "" {
			req.Header.Set("Authorization", "Bearer "+tok)
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, 0, err
		}
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		out := map[string]any{}
		if err := json.Unmarshal(data, &out); err != nil {
			out = map[string]any{"raw": string(data)}
		}
		return out, resp.StatusCode, nil
	}

	out, status, err := send(token)
	if err != nil {
		return nil, err
	}
	// OIDC mode: token expire → refresh sekali lalu retry.
	if status == 401 && c.token == "" && c.clientID != "" {
		c.oidcToken = ""
		c.oidcExpiresAt = 0
		token, err = c.fetchOIDCToken()
		if err != nil {
			return nil, err
		}
		out, status, err = send(token)
		if err != nil {
			return nil, err
		}
	}
	if status >= 400 {
		return nil, fmt.Errorf("TerminologyClient HTTP %d", status)
	}
	return out, nil
}

// resolveToken — PAT statis menang; kalau kosong dan ada OIDC creds → lazy fetch + cache.
func (c *TerminologyClient) resolveToken() (string, error) {
	if c.token != "" {
		return c.token, nil
	}
	if c.clientID == "" {
		return "", nil
	}
	if c.oidcToken != "" && time.Now().Unix() < c.oidcExpiresAt-30 {
		return c.oidcToken, nil
	}
	return c.fetchOIDCToken()
}

// fetchOIDCToken — POST tokenURL grant_type=client_credentials.
func (c *TerminologyClient) fetchOIDCToken() (string, error) {
	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {c.clientID},
		"client_secret": {c.clientSecret},
	}
	resp, err := c.httpClient.PostForm(c.tokenURL, form)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("TerminologyClient OIDC token HTTP %d", resp.StatusCode)
	}
	data, _ := io.ReadAll(resp.Body)
	var tok struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err := json.Unmarshal(data, &tok); err != nil || tok.AccessToken == "" {
		return "", fmt.Errorf("TerminologyClient OIDC token response invalid")
	}
	c.oidcToken = tok.AccessToken
	if tok.ExpiresIn == 0 {
		tok.ExpiresIn = 3600
	}
	c.oidcExpiresAt = time.Now().Unix() + tok.ExpiresIn
	return c.oidcToken, nil
}

// clampLimit — batas 1..100.
func clampLimit(n int) int {
	if n < 1 {
		return 1
	}
	if n > 100 {
		return 100
	}
	return n
}

// Search — GET /api/v1/terminology/search?q=&category=&limit=
func (c *TerminologyClient) Search(q, category string, limit int) (map[string]any, error) {
	if category == "" {
		category = "all"
	}
	qry := url.Values{"q": {q}, "category": {category}, "limit": {strconv.Itoa(clampLimit(limit))}}
	return c.doJSON("GET", "api/v1/terminology/search", qry, nil)
}

// SearchSystem — GET /api/v1/terminology/{system}/search?q=
// system: icd10|icd9cm|snomed|loinc|kfa|kfa_alkes|kfa_bza|kptl
func (c *TerminologyClient) SearchSystem(system, q string, limit int) (map[string]any, error) {
	qry := url.Values{"q": {q}, "limit": {strconv.Itoa(clampLimit(limit))}}
	return c.doJSON("GET", "api/v1/terminology/"+url.PathEscape(system)+"/search", qry, nil)
}

// KptlCombine — GET /api/v1/terminology/kptl/combine?base_code=&modifiers= — multi-layer KPTL.
func (c *TerminologyClient) KptlCombine(baseCode string, modifiers []string) (map[string]any, error) {
	qry := url.Values{"base_code": {baseCode}, "modifiers": {strings.Join(modifiers, ",")}}
	return c.doJSON("GET", "api/v1/terminology/kptl/combine", qry, nil)
}

// Map — GET /api/v1/terminology/mapping?source=&target=&code= — cross-terminology mapping.
func (c *TerminologyClient) Map(source, target, code string) (map[string]any, error) {
	qry := url.Values{"source": {source}, "target": {target}, "code": {code}}
	return c.doJSON("GET", "api/v1/terminology/mapping", qry, nil)
}

// Validate — POST /api/v1/terminology/validate — {system, code}.
func (c *TerminologyClient) Validate(system, code string) (map[string]any, error) {
	return c.doJSON("POST", "api/v1/terminology/validate", nil, map[string]any{"system": system, "code": code})
}

// AutoMap — POST /api/v1/terminology/auto-map — {source_system, source_code} → SNOMED.
func (c *TerminologyClient) AutoMap(sourceSystem, sourceCode string) (map[string]any, error) {
	return c.doJSON("POST", "api/v1/terminology/auto-map", nil, map[string]any{"source_system": sourceSystem, "source_code": sourceCode})
}

// ResolveCondition — POST api/v1/terminology/resolve/condition.
func (c *TerminologyClient) ResolveCondition(code, systemHint, category string) (map[string]any, error) {
	if category == "" {
		category = "encounter-diagnosis"
	}
	return c.resolve("condition", code, systemHint, category)
}

// ResolveChiefComplaint — POST api/v1/terminology/resolve/chief-complaint.
func (c *TerminologyClient) ResolveChiefComplaint(code string) (map[string]any, error) {
	return c.resolve("chief-complaint", code, "", "")
}

// ResolveObservation — POST api/v1/terminology/resolve/observation.
func (c *TerminologyClient) ResolveObservation(code, systemHint string) (map[string]any, error) {
	return c.resolve("observation", code, systemHint, "")
}

// ResolveMedication — POST api/v1/terminology/resolve/medication.
func (c *TerminologyClient) ResolveMedication(code string) (map[string]any, error) {
	return c.resolve("medication", code, "", "")
}

// ResolveServiceRequest — POST api/v1/terminology/resolve/service-request.
func (c *TerminologyClient) ResolveServiceRequest(code, category string) (map[string]any, error) {
	if category == "" {
		category = "lab"
	}
	return c.resolve("service-request", code, "", category)
}

// resolve — POST api/v1/terminology/resolve/{kind}; field kosong di-skip.
func (c *TerminologyClient) resolve(kind, code, systemHint, category string) (map[string]any, error) {
	body := map[string]any{"code": code}
	if systemHint != "" {
		body["system_hint"] = systemHint
	}
	if category != "" {
		body["category"] = category
	}
	return c.doJSON("POST", "api/v1/terminology/resolve/"+url.PathEscape(kind), nil, body)
}

// ValueSets — GET /api/value-sets.
func (c *TerminologyClient) ValueSets() (map[string]any, error) {
	return c.doJSON("GET", "api/value-sets", nil, nil)
}

// ValueSet — GET /api/value-sets/{bundleId}, id di path-escape.
func (c *TerminologyClient) ValueSet(id string) (map[string]any, error) {
	return c.doJSON("GET", "api/value-sets/"+url.PathEscape(id), nil, nil)
}
