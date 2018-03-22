package config

import (
	"testing"
)

func TestIsActive(t *testing.T) {
	c := Config{
		serverList: []Server{new(server)},
		fileReport: "",
	}

	if !c.IsActive() {
		t.Error("Config should be active but it is not")
	}
}

func TestIsNotActive(t *testing.T) {
	c := Config{
		serverList: []Server{},
		fileReport: "",
	}

	if c.IsActive() {
		t.Error("Config should not be active but it is")
	}
}

func TestMakingFileReport(t *testing.T) {
	c := Config{
		serverList: []Server{},
		fileReport: "example.txt",
	}

	if !c.DoFileReport() {
		t.Error("Config should make report in file but it will not")
	}
}

func TestDoNotMakingFileReport(t *testing.T) {
	c := Config{
		serverList: []Server{},
		fileReport: "",
	}

	if c.DoFileReport() {
		t.Error("Config should not make report in file but it will")
	}
}

func TestAddNewServer(t *testing.T) {
	c := Config{
		serverList: []Server{server{url: "http://company.com", disableParams: []string{}}},
		fileReport: "",
	}

	c.AddServer("http://example.com", []string{"param#1"})
	c.AddServer("http://hilda.com", []string{"param#2"})

	if len(c.GetServers()) != 3 {
		t.Errorf("There should be 3 servers in config file, but '%d' are", len(c.GetServers()))
	}
}

func TestAddDuplicatedServer(t *testing.T) {
	c := Config{
		serverList: []Server{server{url: "http://company.com", disableParams: []string{}}},
		fileReport: "",
	}
	c.AddServer("http://example.com", []string{"param#1"})
	c.AddServer("http://company.com", []string{})

	if len(c.GetServers()) != 2 {
		t.Errorf("There should be 2 servers in config file, but are '%d'", len(c.GetServers()))
	}
}
