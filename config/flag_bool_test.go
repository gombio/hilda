package config

import "testing"

func TestShortName(t *testing.T) {
	short := shortName(RequestDisable)

	if short != "request" {
		t.Errorf("Value should be equal to 'request' but '%s' is", short)
	}
}

func TestShortNameForServicesFlag(t *testing.T) {
	short := shortName(ServicesDisable)

	if short != "services" {
		t.Errorf("Value should be equal to 'services' but '%s' is", short)
	}
}

func TestShortNameWithCustomValue(t *testing.T) {
	short := shortName("parameter-val-disable")

	if short != "parameter-val" {
		t.Errorf("Value should be equal to 'parameter-val' but '%s' is", short)
	}
}

func TestShortNameWithoutPostfix(t *testing.T) {
	short := shortName("value")

	if short != "value" {
		t.Errorf("Value should be equal to 'value' but '%s' is", short)
	}
}

func TestLongName(t *testing.T) {
	long := LongName("request")

	if long != RequestDisable {
		t.Errorf("Long value should be equal to 'request-disable' but is: %s", long)
	}
}

func TestLongCustomName(t *testing.T) {
	long := LongName("value-param")

	if long != "value-param-disable" {
		t.Errorf("Long value should be equal to 'value-param-disable' but is: %s", long)
	}
}
