package component

import (
	"net/http"
	"strings"

	ht "github.com/gombio/hilda/test"
)

//Request makes initial requests and test if there are any errors related to them
func Request(client *http.Client, clientTLS *http.Client) Component {
	// if client.Transport.(*http.Transport) {
	//
	// }
	// client := &http.Client{Transport: &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	// }}

	// log.Println(client.Transport)
	// client.Transport.TLSClientConfig.InsecureSkipVerify = true

	return func(c *ht.Context, r *ht.Report) {
		r.Components["request"] = make(map[string]string)

		resp, err := clientTLS.Get(c.URL)
		c.Response = resp
		if err != nil {
			r.StatusError()

			if strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
				r.Components["request"]["TLS"] = err.Error()

				//try client that ignores TLS
				resp, err = client.Get(c.URL)
				c.Response = resp
				if err != nil {
					r.Components["request"]["NO_TLS"] = err.Error()
				}
			} else {
				r.Components["request"]["ERROR"] = err.Error()
			}
		}
	}
}
