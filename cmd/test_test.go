package cmd

import (
	"github.com/gombio/hilda/config"
	"github.com/spf13/cobra"
	"testing"
)

func TestNotShowReport(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().BoolP(config.RequestDisable, "r", true, "")

	if showReport("request", &cmd) {
		t.Error("Flag request is set to true, report should not be displayed")
	}
}

func TestShowReport(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().BoolP(config.RequestDisable, "r", false, "")

	if !showReport("request", &cmd) {
		t.Error("Flag request is set to false, report should be displayed")
	}
}

func TestShowReportWithNoExistFlag(t *testing.T) {
	cmd := cobra.Command{}
	cmd.Flags().BoolP("no exist", "r", false, "")

	if !showReport("request", &cmd) {
		t.Error("Flag request does not exist, report should be displayed")
	}
}
