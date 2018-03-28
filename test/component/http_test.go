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

	HTTP()(ctx, rpt)
	if rpt.Status != "OK" { //TODO: get rid of magic constant
		t.Fatal()
	}

	ctx.Response.StatusCode = 500
	HTTP()(ctx, rpt)
	if rpt.Status != "Error" { //TODO: get rid of magic constant
		t.Fatal()
	}
	if rpt.Components["http"]["status_code"] != "ERROR: Invalid status code 500" {
		t.Fatal()
	}
}
