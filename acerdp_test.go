package acerdp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// The client against a fake API: path parameters, query, body, idempotency key, retries and errors.
func TestClientAgainstAFakeAPI(t *testing.T) {
	tries := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/api/v1/servers" && r.Method == http.MethodGet:
			tries++
			if tries == 1 {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = io.WriteString(w, `{"success":false,"error":{"code":"RATE_LIMITED","message":"slow down"}}`)
				return
			}
			if r.URL.Query().Get("per_page") != "5" || r.Header.Get("Authorization") != "Bearer key" {
				t.Errorf("query %s, auth %s", r.URL.RawQuery, r.Header.Get("Authorization"))
			}
			_, _ = io.WriteString(w, `{"success":true,"data":[{"id":"srv-1"}]}`)
		case r.URL.Path == "/api/v1/servers/srv-1/power" && r.Method == http.MethodPost:
			var body map[string]any
			_ = json.NewDecoder(r.Body).Decode(&body)
			if body["action"] != "restart" || r.Header.Get("Idempotency-Key") != "k1" {
				t.Errorf("body %v, key %q", body, r.Header.Get("Idempotency-Key"))
			}
			_, _ = io.WriteString(w, `{"success":true,"data":{"queued":true}}`)
		default:
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.WriteString(w, `{"success":false,"error":{"code":"RESOURCE_NOT_FOUND","message":"not found"}}`)
		}
	}))
	defer srv.Close()

	var slept time.Duration
	api := &Client{APIKey: "key", BaseURL: srv.URL, HTTP: srv.Client(), Retries: 2, Sleep: func(d time.Duration) { slept += d }}
	ctx := context.Background()

	list, err := api.ListServers(ctx, Params{"per_page": 5})
	if err != nil || string(list) != `[{"id":"srv-1"}]` || slept != time.Second {
		t.Fatalf("list %s, err %v, slept %s", list, err, slept)
	}
	if _, err := api.SendAPowerAction(ctx, Params{"server": "srv-1", "action": "restart", "idempotency_key": "k1"}); err != nil {
		t.Fatal(err)
	}
	_, err = api.GetAServer(ctx, Params{"server": "nope"})
	if apiErr, ok := err.(*Error); !ok || apiErr.Code != "RESOURCE_NOT_FOUND" || apiErr.Status != 404 {
		t.Fatalf("err %v", err)
	}
	if _, err := api.GetAServer(ctx, nil); err == nil {
		t.Fatal("a missing path parameter must be refused")
	}
}
