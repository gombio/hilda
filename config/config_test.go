package config

import "testing"

func TestInitBoolConfigName(t *testing.T) {
	conf := InitBoolConfig("bool-name")

	if conf.cnf.name != "bool-name" {
		t.Errorf("Config name should be equal to 'bool-name' but is: %s", conf.cnf.name)
	}
}

func TestInitBoolConfigDefaultValue(t *testing.T) {
	conf := InitBoolConfig("bool-name")

	if conf.value {
		t.Error("Default value should be equal to 'false'")
	}
}

func TestInitBoolConfigEnvDefault(t *testing.T) {
	conf := InitBoolConfig("bool-name")

	if conf.envValue {
		t.Error("Config env value should not be set after init")
	}

	if conf.cnf.valueFromEnv {
		t.Error("Config should inform value from env was not setup aftr init")
	}
}

func TestInitBoolConfigFileDefault(t *testing.T) {
	conf := InitBoolConfig("bool-name")

	if conf.fileValue {
		t.Error("Config file value should not be set after init")
	}
	if conf.cnf.valueFromFile {
		t.Error("Config should inform value from file was not setup aftr init")
	}
}

func TestSetFileValue(t *testing.T) {
	conf := InitBoolConfig("bool-name")
	conf.SetFileValue(true)

	if !conf.fileValue {
		t.Error("Config file value should be equal to 'true' but it is not")
	}
	if !conf.cnf.valueFromFile {
		t.Error("Config should inform value from file was setup but it is not")
	}
}

func TestSetEnvValue(t *testing.T) {
	conf := InitBoolConfig("bool-name")
	conf.SetEnvValue(true)

	if !conf.envValue {
		t.Error("Config env value should be equal to 'true' but it is not")
	}
	if !conf.cnf.valueFromEnv {
		t.Error("Config should inform value from env was setup but it is not")
	}
}

func TestGetWithEnvValue(t *testing.T) {
	conf := InitBoolConfig("bool-name")
	conf.SetFileValue(true)
	conf.SetEnvValue(false)

	if conf.Get() {
		t.Error("Value should be taken from env")
	}
}

func TestGetWithFileValue(t *testing.T) {
	conf := InitBoolConfig("bool-name")
	conf.SetFileValue(true)

	if !conf.Get() {
		t.Error("Value should be taken from file")
	}
}

func TestGetWithDefaultValue(t *testing.T) {
	conf := InitBoolConfig("bool-name")

	if conf.Get() {
		t.Error("Value should be default equal 'false'")
	}
}
