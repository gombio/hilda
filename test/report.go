package test

const StatusOk string = "OK"
const StatusWarning string = "Warning"
const StatusError string = "Error"

type Report struct {
	URL        string
	Status     string
	Components map[string]map[string]string
}

func NewReport(url string) *Report {
	return &Report{
		URL:        url,
		Status:     StatusOk,
		Components: make(map[string]map[string]string),
	}
}

func (r *Report) StatusOk() {
	r.Status = StatusOk
}

func (r *Report) StatusWarning() {
	r.Status = StatusWarning
}

func (r *Report) StatusError() {
	r.Status = StatusError
}
