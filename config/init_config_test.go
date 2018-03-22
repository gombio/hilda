package config

import "testing"

var cf = configFile{
	Server: map[string]reportConfig{
		"http://example.com": reportConfig{
			Params: map[string]string{"ca": "disabled"},
		},
	},
	ReportFile: "debug.txt",
}

func TestConfigFileReport(t *testing.T) {

	var config Config
	config.fileReport = cf.ReportFile

	if !config.DoFileReport() {
		t.Error("Config should indicate to create file report but it didn't")
	}
	if config.fileReport != "debug.txt" {
		t.Error("Config has wrond file report data")
	}
}

func TestConfigServer(t *testing.T) {

	var config Config
	config.serverList = initServers(cf.Server)

	if !config.IsActive() {
		t.Error("Config should be active if has at least one server")
	}
}

func TestConfigServerWithWrongData(t *testing.T) {
	cf := configFile{
		Server: map[string]reportConfig{
			"wrong url": reportConfig{
				Params: map[string]string{"ca": "disabled"},
			},
		},
		ReportFile: "",
	}

	var config Config
	config.serverList = initServers(cf.Server)

	if config.IsActive() {
		t.Error("Config should not be active beacue init data is wrong")
	}
}
