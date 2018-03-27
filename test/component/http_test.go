package component

import (
	"testing"

	ht "github.com/gombio/hilda/test"
)

func TestHttp(t *testing.T) {
	url := "test"
	ctx := ht.NewContext(url)
	ctx.Response.StatusCode = 200
	rpt := ht.NewReport(url)

	Http()(ctx, rpt)
	if rpt.Status != ht.StatusOk {
		t.Fatal()
	}

	ctx.Response.StatusCode = 500
	Http()(ctx, rpt)
	if rpt.Status != ht.StatusError {
		t.Fatal()
	}
	if rpt.Components["http"]["status_code"] != "ERROR: Invalid status code 500" {
		t.Fatal()
	}
}
