package config

import "strings"

const suffix = "-disable"

// RequestDisable flag name to disable request from test report
const RequestDisable = "request-disable"

// HTTPDisable flag name to disable http from test report
const HTTPDisable = "http-disable"

// ServicesDisable flag name to disable services from test report
const ServicesDisable = "services-disable"

type boolFlag struct {
	name         string
	short        string
	defaultValue bool
	description  string
}

var boolFlags = []boolFlag{
	boolFlag{
		name:         RequestDisable,
		short:        "r",
		defaultValue: false,
		description:  "Disable check server parameter",
	},
	boolFlag{
		name:         HTTPDisable,
		short:        "w",
		defaultValue: false,
		description:  "Disable check http parameter",
	},
	boolFlag{
		name:         ServicesDisable,
		short:        "s",
		defaultValue: false,
		description:  "Disable check services parameter",
	},
}

func shortName(name string) string {
	return strings.TrimSuffix(name, suffix)
}

// LongName add suffix to short flag name
// to be consist with bool flag name
func LongName(name string) string {
	return name + suffix
}
