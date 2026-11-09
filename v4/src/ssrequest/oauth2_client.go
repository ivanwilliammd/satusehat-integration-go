package ssrequest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// Default sandbox test IDs from SATUSEHAT.
var (
	DefaultPatientDev      = []string{"P02478375538", "P02428473601", "P03647103112", "P01058967035", "P01836748436", "P01654557057", "P00805884304", "P00883356749", "P00912894463"}
	DefaultPractitionerDev = []string{"10009880728", "10006926841", "10001354453", "10010910332", "10018180913", "10002074224", "10012572188", "10018452434", "10014058550", "10001915884"}
)

// OAuth2Config holds optional configuration overrides for OAuth2Client.
type OAuth2Config struct {
	Env            string
	BaseURL        string
	AuthEndpoint   string
	FhirEndpoint   string
	ClientID       string
	ClientSecret   string
	OrganizationID string
	Override       bool
	AutoRefresh    *bool
	HTTPClient     *http.Client
}

// OAuth2Client manages SATUSEHAT OAuth2 authentication and direct helper endpoints.
type OAuth2Client struct {
	PatientDev      []string
	PractitionerDev []string
	BaseURL         string
	AuthURL         string
	FhirURL         string
	ClientID        string
	ClientSecret    string
	OrganizationID  string
	Override        bool
	SatusehatEnv    string
	AutoRefresh     bool
	HTTPClient      *http.Client

	mu          sync.Mutex
	cachedToken string
	tokenExpiry time.Time
}

// NewOAuth2Client initializes an OAuth2Client from OAuth2Config or environment variables.
func NewOAuth2Client(configs ...OAuth2Config) (*OAuth2Client, error) {
	var cfg OAuth2Config
	if len(configs) > 0 {
		cfg = configs[0]
	}

	autoRefresh := true
	if cfg.AutoRefresh != nil {
		autoRefresh = *cfg.AutoRefresh
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}

	c := &OAuth2Client{
		PatientDev:      append([]string(nil), DefaultPatientDev...),
		PractitionerDev: append([]string(nil), DefaultPractitionerDev...),
		Override:        cfg.Override,
		AutoRefresh:     autoRefresh,
		HTTPClient:      httpClient,
	}

	if !c.Override {
		env := cfg.Env
		if env == "" {
			env = os.Getenv("SATUSEHAT_ENV")
		}
		c.SatusehatEnv = env

		switch c.SatusehatEnv {
		case "PROD":
			c.BaseURL = firstNonEmpty(cfg.BaseURL, os.Getenv("SATUSEHAT_BASE_URL_PROD"), "https://api-satusehat.kemkes.go.id")
			c.ClientID = firstNonEmpty(cfg.ClientID, os.Getenv("CLIENTID_PROD"))
			c.ClientSecret = firstNonEmpty(cfg.ClientSecret, os.Getenv("CLIENTSECRET_PROD"))
			c.OrganizationID = firstNonEmpty(cfg.OrganizationID, os.Getenv("ORGID_PROD"))
		case "STG":
			c.BaseURL = firstNonEmpty(cfg.BaseURL, os.Getenv("SATUSEHAT_BASE_URL_STG"), "https://api-satusehat-stg.dto.kemkes.go.id")
			c.ClientID = firstNonEmpty(cfg.ClientID, os.Getenv("CLIENTID_STG"))
			c.ClientSecret = firstNonEmpty(cfg.ClientSecret, os.Getenv("CLIENTSECRET_STG"))
			c.OrganizationID = firstNonEmpty(cfg.OrganizationID, os.Getenv("ORGID_STG"))
		case "DEV":
			c.BaseURL = firstNonEmpty(cfg.BaseURL, os.Getenv("SATUSEHAT_BASE_URL_DEV"), "https://api-satusehat-dev.dto.kemkes.go.id")
			c.ClientID = firstNonEmpty(cfg.ClientID, os.Getenv("CLIENTID_DEV"))
			c.ClientSecret = firstNonEmpty(cfg.ClientSecret, os.Getenv("CLIENTSECRET_DEV"))
			c.OrganizationID = firstNonEmpty(cfg.OrganizationID, os.Getenv("ORGID_DEV"))
		}

		if c.SatusehatEnv == "" {
			return nil, fmt.Errorf("SATUSEHAT environment is missing")
		}
		if c.SatusehatEnv != "DEV" && c.SatusehatEnv != "STG" && c.SatusehatEnv != "PROD" {
			return nil, fmt.Errorf("SATUSEHAT environment invalid, supported (DEV, STG, PROD). %s given.", c.SatusehatEnv)
		}
		if c.ClientID == "" || c.ClientSecret == "" || c.OrganizationID == "" {
			return nil, fmt.Errorf("SATUSEHAT environment defined as %s, but CLIENTID_%s / CLIENTSECRET_%s / ORGID_%s not set",
				c.SatusehatEnv, c.SatusehatEnv, c.SatusehatEnv, c.SatusehatEnv)
		}
	} else {
		c.SatusehatEnv = cfg.Env
		c.BaseURL = cfg.BaseURL
		c.ClientID = cfg.ClientID
		c.ClientSecret = cfg.ClientSecret
		c.OrganizationID = cfg.OrganizationID
	}

	authEndpoint := firstNonEmpty(cfg.AuthEndpoint, os.Getenv("SATUSEHAT_AUTH_ENDPOINT"), "/oauth2/v1")
	fhirEndpoint := firstNonEmpty(cfg.FhirEndpoint, os.Getenv("SATUSEHAT_FHIR_ENDPOINT"), "/fhir-r4/v1")

	c.BaseURL = strings.TrimRight(c.BaseURL, "/")
	c.AuthURL = c.BaseURL + authEndpoint
	c.FhirURL = c.BaseURL + fhirEndpoint

	return c, nil
}

// IsTokenAutoRefresh returns whether automatic token refresh on 401 is enabled.
func (c *OAuth2Client) IsTokenAutoRefresh() bool {
	return c.AutoRefresh
}

// Token returns a valid OAuth2 access token, using a 50-minute cache when available.
func (c *OAuth2Client) Token() (string, error) {
	c.mu.Lock()
	if c.cachedToken != "" && time.Now().Before(c.tokenExpiry) {
		tok := c.cachedToken
		c.mu.Unlock()
		return tok, nil
	}
	c.mu.Unlock()

	return c.RefreshToken()
}

// RefreshToken forces fetching a new OAuth2 access token from SATUSEHAT.
func (c *OAuth2Client) RefreshToken() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cachedToken = ""
	c.tokenExpiry = time.Time{}

	form := url.Values{}
	form.Set("client_id", c.ClientID)
	form.Set("client_secret", c.ClientSecret)

	tokenURL := strings.TrimRight(c.AuthURL, "/") + "/accesstoken?grant_type=client_credentials"
	req, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rawBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var parsed map[string]any
	_ = json.Unmarshal(rawBytes, &parsed)

	if resp.StatusCode >= 400 {
		issueText := extractOAuthIssue(parsed)
		if issueText == "" {
			issueText = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(rawBytes))
		}
		return "", fmt.Errorf("%s", issueText)
	}

	accessToken, _ := parsed["access_token"].(string)
	if accessToken == "" {
		return "", fmt.Errorf("Unauthorized. Token not found")
	}

	c.cachedToken = accessToken
	c.tokenExpiry = time.Now().Add(50 * time.Minute)
	return accessToken, nil
}

// GetByID fetches a FHIR resource by its type and ID.
func (c *OAuth2Client) GetByID(resource, id string) (int, map[string]any, error) {
	accessToken, err := c.Token()
	if err != nil || accessToken == "" {
		return 401, map[string]any{"error": "Unauthorized. Token not found"}, err
	}

	targetURL := strings.TrimRight(c.FhirURL, "/") + "/" + strings.TrimLeft(resource, "/") + "/" + id
	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	return c.executeJSON(req)
}

// GetByNIK queries a FHIR resource (e.g. Patient or Practitioner) by Indonesian NIK.
func (c *OAuth2Client) GetByNIK(resource, nik string) (int, map[string]any, error) {
	accessToken, err := c.Token()
	if err != nil || accessToken == "" {
		return 401, map[string]any{"error": "Unauthorized. Token not found"}, err
	}

	targetURL := strings.TrimRight(c.FhirURL, "/") + "/" + strings.TrimLeft(resource, "/") + "?identifier=https://fhir.kemkes.go.id/id/nik|" + nik
	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	return c.executeJSON(req)
}

// KfaGet sends a GET request to the SATUSEHAT KFA v2 master data endpoint.
func (c *OAuth2Client) KfaGet(resource, queryString string) (int, map[string]any, error) {
	accessToken, err := c.Token()
	if err != nil || accessToken == "" {
		return 401, map[string]any{"error": "Unauthorized. Token not found"}, err
	}

	targetURL := strings.TrimRight(c.BaseURL, "/") + "/kfa-v2/" + strings.TrimLeft(resource, "/") + queryString
	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	return c.executeJSON(req)
}

// Post sends a raw JSON body to a FHIR resource endpoint (or Bundle root when resource == "Bundle").
func (c *OAuth2Client) Post(resource, body string) (int, map[string]any, error) {
	accessToken, err := c.Token()
	if err != nil || accessToken == "" {
		return 401, map[string]any{"error": "Unauthorized. Token not found"}, err
	}

	targetURL := strings.TrimRight(c.FhirURL, "/")
	if resource != "Bundle" && resource != "" {
		targetURL += "/" + strings.TrimLeft(resource, "/")
	}

	req, err := http.NewRequest(http.MethodPost, targetURL, strings.NewReader(body))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	return c.executeJSON(req)
}

// Put sends a raw JSON body via HTTP PUT to update a FHIR resource by ID.
func (c *OAuth2Client) Put(resource, id, body string) (int, map[string]any, error) {
	accessToken, err := c.Token()
	if err != nil || accessToken == "" {
		return 401, map[string]any{"error": "Unauthorized. Token not found"}, err
	}

	targetURL := strings.TrimRight(c.FhirURL, "/") + "/" + strings.TrimLeft(resource, "/") + "/" + id
	req, err := http.NewRequest(http.MethodPut, targetURL, strings.NewReader(body))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	return c.executeJSON(req)
}

func (c *OAuth2Client) executeJSON(req *http.Request) (int, map[string]any, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	rawBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	var parsed map[string]any
	if len(rawBytes) > 0 {
		_ = json.Unmarshal(rawBytes, &parsed)
	}
	if parsed == nil {
		parsed = map[string]any{}
	}

	return resp.StatusCode, parsed, nil
}

func extractOAuthIssue(parsed map[string]any) string {
	if parsed == nil {
		return ""
	}
	if issues, ok := parsed["issue"].([]any); ok && len(issues) > 0 {
		if first, ok := issues[0].(map[string]any); ok {
			if details, ok := first["details"].(map[string]any); ok {
				if text, ok := details["text"].(string); ok {
					return text
				}
			}
		}
	}
	return ""
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
