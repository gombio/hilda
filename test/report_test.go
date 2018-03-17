package test

import "testing"

func TestNewReport(t *testing.T) {
	rpt := NewReport("foo")
	if rpt.URL != "foo" {
		t.Fatal("Incorrect URL")
	}
}

func TestStatusOk(t *testing.T) {
	rpt := NewReport("foo")
	rpt.Status = "INCORRECT"
	rpt.StatusOk()
	if rpt.Status != StatusOk {
		t.Fatal("Incorrect status")
	}
}

func TestStatusWarning(t *testing.T) {
	rpt := NewReport("foo")
	rpt.Status = "INCORRECT"
	rpt.StatusWarning()
	if rpt.Status != StatusWarning {
		t.Fatal("Incorrect status")
	}
}

func TestStatusError(t *testing.T) {
	rpt := NewReport("foo")
	rpt.Status = "INCORRECT"
	rpt.StatusError()
	if rpt.Status != StatusError {
		t.Fatal("Incorrect status")
	}
}
