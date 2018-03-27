package component

import (
	"crypto/tls"
	"net/http"
	"testing"

	ht "github.com/gombio/hilda/test"
)

//TODO: Refactor so that we can mock http
func TestRequest(t *testing.T) {
	url := "test"
	ctx := ht.NewContext(url)
	rpt := ht.NewReport(url)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	clientTLS := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}}

	Request(client, clientTLS)(ctx, rpt)
}
