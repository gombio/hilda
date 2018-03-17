package test

import (
	"net/http"
)

type Context struct {
	URL      string
	Response *http.Response
}

func NewContext(url string) *Context {
	return &Context{
		URL:      url,
		Response: &http.Response{},
	}
}
