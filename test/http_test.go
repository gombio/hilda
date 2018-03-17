package test

import "testing"

func TestHttp(t *testing.T) {
	url := "test"
	ctx := NewContext(url)
	ctx.Response.StatusCode = 200
	rpt := NewReport(url)

	Http(ctx, rpt)
	if rpt.Status != StatusOk {
		t.Fatal()
	}

	ctx.Response.StatusCode = 500
	Http(ctx, rpt)
	if rpt.Status != StatusError {
		t.Fatal()
	}
	if rpt.Components["http"]["status_code"] != "ERROR: Invalid status code 500" {
		t.Fatal()
	}
}
