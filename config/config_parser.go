package config

type configFile struct {
	Server     map[string]reportConfig `yaml:"server"`
	ReportFile string                  `yaml:"report_file"`
}

type reportConfig struct {
	Params map[string]string
}

// get list of disabled parameters
func (rc reportConfig) get() []string {
	var result []string
	for key, value := range rc.Params {
		if value == "disable" {
			result = append(result, key)
		}
	}
	return result
}

func (rc *reportConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {

	rc.Params = make(map[string]string)
	var multi map[string]string

	err := unmarshal(&multi)
	if err == nil {
		for key, value := range multi {
			rc.Params[key] = value
		}
	}

	return nil
}
