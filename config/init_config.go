package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func Init(file string) *Config {
	config, err := initConfig(file)
	if err != nil {
		log.Printf("Config file '%s' cannot be create (%s) \n", file, err)
	}

	return config
}

func initConfig(file string) (*Config, error) {

	configFile := configFile{}

	err := yaml.Unmarshal(readFile(file), &configFile)
	if err != nil {
		return &Config{}, err
	}

	config := new(Config)
	config.serverList = initServers(configFile.Server)
	config.fileReport = configFile.ReportFile

	return config, nil
}

// init servers struct; server url + disable params to report
func initServers(servers map[string]reportConfig) []Server {
	var serverList []Server
	for url, params := range servers {
		server, err := createServer(url, params.get())
		if err != nil {
			log.Println(err)
			continue
		}
		serverList = append(serverList, server)
	}

	return serverList
}

// read yaml file with configuration
func readFile(file string) []byte {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Unable to read config file: %v \n", err)
	}

	return []byte(yamlFile)
}
