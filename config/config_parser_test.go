package config

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

var fileContent = `
server:
    https://example.com:
        ca: 'disable'
    https://global-corm.com: ~

report_file: 'debug.txt'
`

func TestUnmarshalConfigFile(t *testing.T) {
	cf := configFile{}

	err := yaml.Unmarshal([]byte(fileContent), &cf)
	if err != nil {
		t.Errorf("Yaml file cannot be unmarshal: %s", err)
	}
}

func TestReportFileParameter(t *testing.T) {
	cf := configFile{}
	yaml.Unmarshal([]byte(fileContent), &cf)

	if cf.ReportFile != "debug.txt" {
		t.Error("Report file parameter is incorrect")
	}
}

func TestServerParameters(t *testing.T) {
	cf := configFile{}
	yaml.Unmarshal([]byte(fileContent), &cf)

	for url, params := range cf.Server {
		if url == "https://example.com" {
			eq := reflect.DeepEqual(params, reportConfig{Params: map[string]string{"ca": "disable"}})
			if !eq {
				t.Error("Servers parameters are incorrect for specific server")
			}
			continue
		}
		if url == "https://global-corm.com" {
			continue
		}
		t.Error("Server parameters are incorrect")
	}
}
