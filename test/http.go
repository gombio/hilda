package test

import (
	"fmt"
	"strconv"
)

func Http(c *Context, r *Report) {
	r.Components["http"] = make(map[string]string)

	//if http status is not 2xx - make error
	if c.Response.StatusCode < 200 || 299 < c.Response.StatusCode {
		r.StatusError()
		r.Components["http"]["status_code"] = fmt.Sprintf("ERROR: Invalid status code %s", strconv.Itoa(c.Response.StatusCode))
	}
}
