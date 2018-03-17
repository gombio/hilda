package test

import "testing"

//TODO: Refactor so that we can mock http
func TestRequest(t *testing.T) {
	url := "test"
	ctx := NewContext(url)
	rpt := NewReport(url)

	Request(ctx, rpt)
}
