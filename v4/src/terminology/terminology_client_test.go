package terminology

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// newTestClient — client dgn baseURL menunjuk ke server tes.
func newTestClient(t *testing.T, srvURL string) *TerminologyClient {
	t.Helper()
	c := NewTerminologyClient(srvURL, "pat-test")
	return c
}

func TestSearchSystem(t *testing.T) {
	var gotPath, gotQ string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQ = r.URL.RawQuery
		if r.Header.Get("Authorization") != "Bearer pat-test" {
			t.Errorf("authorization header = %q, want Bearer pat-test", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"code":"1234","display":"Tes"}]}`))
	}))
	defer srv.Close()

	res, err := newTestClient(t, srv.URL).SearchSystem("snomed", "demam", 5)
	if err != nil {
		t.Fatalf("SearchSystem error: %v", err)
	}
	if gotPath != "/api/v1/terminology/snomed/search" {
		t.Errorf("path = %q", gotPath)
	}
	if !strings.Contains(gotQ, "q=demam") || !strings.Contains(gotQ, "limit=5") {
		t.Errorf("query = %q, want q=demam&limit=5", gotQ)
	}
	if res["data"] == nil {
		t.Errorf("body tidak ter-decode: %v", res)
	}
}

func TestQueryBuilderChaining(t *testing.T) {
	var gotPath, gotQ string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQ = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := newTestClient(t, srv.URL)
	// Query().On().Category().Limit().Get() → SearchSystem (On menang atas category)
	if _, err := c.Query("demam").On("kfa").Category("diagnosis").Limit(3).Get(); err != nil {
		t.Fatalf("chaining error: %v", err)
	}
	if gotPath != "/api/v1/terminology/kfa/search" || !strings.Contains(gotQ, "q=demam") || !strings.Contains(gotQ, "limit=3") {
		t.Errorf("search-system hit = %s?%s", gotPath, gotQ)
	}

	// KptlCombine chaining → GET /kptl/combine dgn base_code + modifiers gabung koma
	if _, err := c.KptlCombine("1.02.01", []string{"1.02.01.01", "1.02.01.02"}); err != nil {
		t.Fatalf("KptlCombine error: %v", err)
	}
	if gotPath != "/api/v1/terminology/kptl/combine" {
		t.Errorf("kptl path = %q", gotPath)
	}
	if !strings.Contains(gotQ, "base_code=1.02.01") || !strings.Contains(gotQ, "modifiers=1.02.01.01%2C1.02.01.02") {
		t.Errorf("kptl query = %q", gotQ)
	}
}

func TestAsCodeGuessSystem(t *testing.T) {
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/api/v1/terminology/validate" {
			t.Errorf("hit = %s %s", r.Method, r.URL.Path)
		}
		json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"valid":true}`))
	}))
	defer srv.Close()

	if _, err := newTestClient(t, srv.URL).Query("93004944").AsCode().Get(); err != nil {
		t.Fatalf("AsCode error: %v", err)
	}
	if gotBody["system"] != "kfa" {
		t.Errorf("system = %v, want kfa (tebakan dari prefix 93xxxx)", gotBody["system"])
	}
	if gotBody["code"] != "93004944" {
		t.Errorf("code = %v", gotBody["code"])
	}
}

func TestHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"unauthorized"}`))
	}))
	defer srv.Close()

	_, err := newTestClient(t, srv.URL).Search("demam", "all", 20)
	if err == nil || !strings.Contains(err.Error(), "TerminologyClient HTTP 401") {
		t.Fatalf("err = %v, want TerminologyClient HTTP 401", err)
	}
}

func TestValidateStaticTokenWins(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer pat" {
			t.Errorf("Authorization = %q, want Bearer pat", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := NewTerminologyClient(srv.URL, "pat")
	if _, err := c.Validate("kfa", "93004944"); err != nil {
		t.Fatalf("Validate error: %v", err)
	}
}

func TestOIDCLazyFetch(t *testing.T) {
	var authHits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/oauth/token":
			if r.Method != "POST" || r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				t.Errorf("token hit = %s %s", r.Method, r.Header.Get("Content-Type"))
			}
			r.ParseForm()
			if r.Form.Get("grant_type") != "client_credentials" ||
				r.Form.Get("client_id") != "cid" ||
				r.Form.Get("client_secret") != "csecret" {
				t.Errorf("token form = %v", r.Form)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"access_token":"jwt-abc","expires_in":3600}`))
		default:
			// hit 1: 401 → refresh-once lalu retry; hit 2: token cached valid.
			n := atomic.AddInt32(&authHits, 1)
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"message":"expired"}`))
				return
			}
			if got := r.Header.Get("Authorization"); got != "Bearer jwt-abc" {
				t.Errorf("Authorization = %q, want Bearer jwt-abc", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"ok":true}`))
		}
	}))
	defer srv.Close()

	c := NewTerminologyClient(srv.URL, "")
	c.clientID = "cid"
	c.clientSecret = "csecret"
	c.tokenURL = srv.URL + "/oauth/token"
	t.Setenv("TERMINOLOGY_API_TOKEN", "")
	t.Setenv("TERMINOLOGY_CLIENT_ID", "cid")
	t.Setenv("TERMINOLOGY_CLIENT_SECRET", "csecret")
	t.Setenv("TERMINOLOGY_TOKEN_URL", srv.URL+"/oauth/token")

	if _, err := c.Search("demam", "all", 10); err != nil {
		t.Fatalf("OIDC Search error: %v", err)
	}
	if _, err := c.Search("demam", "all", 10); err != nil {
		t.Fatalf("OIDC cached Search error: %v", err)
	}
	if atomic.LoadInt32(&authHits) != 3 {
		t.Errorf("api hits = %d (401 + retry + cached), want 3", atomic.LoadInt32(&authHits))
	}
}
