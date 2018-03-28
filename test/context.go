package test

import (
	"net/http"
)

//Context carries context of execution accross various tests
type Context struct {
	URL      string
	Response *http.Response
}

//NewContext creates new test context
func NewContext(url string) *Context {
	return &Context{
		URL:      url,
		Response: &http.Response{},
	}
}
