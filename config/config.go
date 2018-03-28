package config

import (
	"log"
	"strings"
)

//Configuration TODO: description
type Configuration interface {
	IsActive() bool
	DoFileReport() bool
	GetServers() []Server
	AddServer(url string, params map[string]string)
}

//Config TODO: description
type Config struct {
	serverList []Server
	fileReport string
}

//IsActive TODO: description
func (c Config) IsActive() bool {
	return len(c.serverList) > 0
}

//DoFileReport TODO: description
func (c Config) DoFileReport() bool {
	return len(c.fileReport) > 0
}

//GetServers TODO: description
func (c Config) GetServers() []Server {
	return c.serverList
}

//AddServer TODO: description
func (c *Config) AddServer(url string, params []string) {

	for _, server := range c.serverList {
		if strings.Contains(server.GetURL(), url) {
			return
		}
	}

	server, err := createServer(url, params)
	//TODO: provide tests
	if err != nil {
		log.Printf("Server with url: '%s' has not been added \n", url)
		return
	}
	c.serverList = append(c.serverList, server)
}
