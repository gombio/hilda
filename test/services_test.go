package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServices(t *testing.T) {
	url := "http://example.com/healthz"
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "THIS IS NOT A JSON")
	}
	handler(w, req)

	ctx := NewContext(url)
	ctx.Response = w.Result()
	rpt := NewReport(url)
	Services(ctx, rpt)
	if rpt.Status != StatusError {
		t.Fatal("non-JSON input should be marked as error")
	}
	if rpt.Components["services"]["error"] != "Error decoding JSON: invalid character 'T' looking for beginning of value" {
		t.Fatal("non-JSON input should produce JSON validation error message")
	}

	//One service down
	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()

	handler = func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"redis\":\"Error\",\"build\":\"tag: 123 build: 381\",\"database\":\"OK\"}")
	}
	handler(w, req)

	ctx = NewContext(url)
	ctx.Response = w.Result()
	rpt = NewReport(url)
	Services(ctx, rpt)

	if rpt.Status != StatusError {
		t.Fatal("If any service is down, whole report should have status Error")
	}
	if rpt.Components["services"]["redis"] != StatusError {
		t.Fatal("expected failed redis service")
	}

	//All services OK
	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()

	handler = func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"redis\":\"OK\",\"build\":\"tag: 123 build: 381\",\"database\":\"OK\"}")
	}
	handler(w, req)

	ctx = NewContext(url)
	ctx.Response = w.Result()
	rpt = NewReport(url)
	Services(ctx, rpt)
	if rpt.Status != StatusOk {
		t.Fatal("If all servives are ok, report should have StatusOk")
	}
	for svc, status := range rpt.Components["services"] {
		if status != StatusOk {
			t.Fatalf("Service %s should have status OK", svc)
		}
	}
}
