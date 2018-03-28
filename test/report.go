package test

const statusOk string = "OK"
const statusWarning string = "Warning"
const statusError string = "Error"

//Report keeps test results in one place
type Report struct {
	URL        string
	Status     string
	Components map[string]map[string]string
}

//NewReport creates Report instance for provided url
func NewReport(url string) *Report {
	return &Report{
		URL:        url,
		Status:     statusOk,
		Components: make(map[string]map[string]string),
	}
}

//StatusOk sets report status to OK
func (r *Report) StatusOk() {
	r.Status = statusOk
}

//StatusWarning sets report status to Warning
func (r *Report) StatusWarning() {
	r.Status = statusWarning
}

//StatusError sets report status to Error
func (r *Report) StatusError() {
	r.Status = statusError
}
