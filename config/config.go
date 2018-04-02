package config

type config struct {
	name          string
	valueFromFile bool
	valueFromEnv  bool
}

// additional configuration from file/env if flag is not set
type BoolConfig struct {
	cnf       config
	fileValue bool
	envValue  bool
	value     bool
}

// init BoolConfig type with default values
func InitBoolConfig(name string) *BoolConfig {
	return &BoolConfig{
		cnf: config{
			name:          name,
			valueFromFile: false,
			valueFromEnv:  false,
		},
		fileValue: false,
		envValue:  false,
		value:     false,
	}
}

// SetFileValue set value from config file
func (cb *BoolConfig) SetFileValue(val bool) {
	cb.fileValue = val
	cb.cnf.valueFromFile = true
}

// SetEnvValue set value from env variable
func (cb *BoolConfig) SetEnvValue(val bool) {
	cb.envValue = val
	cb.cnf.valueFromEnv = true
}

// Get return proper value depends on hierarchy of the validity of the flag
// if env value is set ten return env value
// if env is not set return file value
// if non of them is set return default value
func (cb BoolConfig) Get() bool {
	if cb.cnf.valueFromEnv {
		return cb.envValue
	}
	if cb.cnf.valueFromFile {
		return cb.fileValue
	}
	return cb.value
}
