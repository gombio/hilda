package verify

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

const StatusOK string = "OK"
const StatusWARNING string = "Warning"
const StatusERROR string = "Error"

type Report struct {
	URL            string
	Status         string //OK, WARNING, ERROR
	ErrorMessage   string //ERROR: error status
	ServicesStatus map[string]string
	SSLStatus      string //OK, WARNING: Warning message, ERROR: Error message
}

func verify(url string) Report {
	trSSL := &http.Transport{}
	trSSLOff := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clientSSL := &http.Client{Transport: trSSL}
	clientSSLOff := &http.Client{Transport: trSSLOff}

	r := Report{
		URL:       url,
		Status:    StatusOK,
		SSLStatus: StatusOK,
	}
	resp, err := clientSSL.Get(url)
	if err != nil {
		tmpError := err.Error()
		r.Status = StatusERROR
		r.ErrorMessage = tmpError

		//test if the error is SSL related
		if strings.Contains(tmpError, "x509: certificate signed by unknown authority") {
			resp, err = clientSSLOff.Get(url)
			if err != nil {
				r.Status = StatusERROR
				r.ErrorMessage = err.Error() //something's wrong with the URL
				fmt.Println("ERROR")

				return r
			}

			// we have recovered, so this is a warning
			r.Status = StatusWARNING
			r.ErrorMessage = ""
			r.SSLStatus = "ERROR: " + tmpError //something's wrong with SSL
		}
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// var services map[string]string
	err = json.Unmarshal(body, &r.ServicesStatus)
	if err != nil {
		r.Status = StatusERROR
		r.ErrorMessage = "Could not decode response - probably not a JSON"

		return r
	}
	//Look for values other then OK in services
	for service, status := range r.ServicesStatus {
		if "build" != service && "OK" != status {
			r.Status = StatusERROR
			r.ErrorMessage = "Services broken"
		}
	}

	return r
}

func Verify(url string) {
	r := verify(url)
	// var cInfo = color.New(color.Bold, color.FgGreen).PrintlnFunc()
	// var cError = color.New(color.Bold, color.FgRed).PrintlnFunc()
	// var cWarning = color.New(color.Bold, color.FgYellow).PrintlnFunc()
	cPrint := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	switch r.Status {
	case StatusOK:
		break
	case StatusWARNING:
		cPrint = color.New(color.Bold, color.FgYellow).PrintlnFunc()
		break
	case StatusERROR:
		cPrint = color.New(color.Bold, color.FgRed).PrintlnFunc()
		break
	}

	fmt.Println("")
	cPrint(r.URL)
	cPrint("=> Status: " + r.Status)
	if r.ErrorMessage != "" {
		cPrint("=> ErrorMessage: " + r.ErrorMessage)
	}
	cPrint("=> SSLStatus: " + r.SSLStatus)
	if len(r.ServicesStatus) > 0 {
		cPrint("=> Services:")
		for service, status := range r.ServicesStatus {
			cPrint("--> " + service + ": " + status)
		}
	}
}
