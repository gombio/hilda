package component

import (
	"fmt"
	"strconv"

	ht "github.com/gombio/hilda/test"
)

func Http() Component {
	return func(c *ht.Context, r *ht.Report) {
		r.Components["http"] = make(map[string]string)

		//if http status is not 2xx - make error
		if c.Response.StatusCode < 200 || 299 < c.Response.StatusCode {
			r.StatusError()
			r.Components["http"]["status_code"] = fmt.Sprintf("ERROR: Invalid status code %s", strconv.Itoa(c.Response.StatusCode))
		}
	}
}
