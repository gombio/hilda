package test

import (
	"encoding/json"
	"io/ioutil"
)

func Services(c *Context, r *Report) {
	r.Components["services"] = make(map[string]string)

	var serviceStatus map[string]string
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		r.StatusError()
		r.Components["services"]["error"] = "Error reading response body: " + err.Error()

		return
	}

	err = json.Unmarshal(body, &serviceStatus)
	if err != nil {
		r.StatusError()
		r.Components["services"]["error"] = "Error decoding JSON: " + err.Error()

		return
	}

	//Look for values other then OK in services
	for service, status := range serviceStatus {
		if "build" != service && "OK" != status {
			r.StatusError()
			r.Components["services"][service] = status
		}
	}
}
