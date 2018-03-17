package test

import (
	"crypto/tls"
	"net/http"
	"strings"
)

func Request(c *Context, r *Report) {
	r.Components["request"] = make(map[string]string)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}}

	resp, err := client.Get(c.URL)
	c.Response = resp
	if err != nil {
		r.StatusError()

		if strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
			r.Components["request"]["TLS"] = err.Error()

			//try ignoring TLS
			client = &http.Client{Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}}
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
