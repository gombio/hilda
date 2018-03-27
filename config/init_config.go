package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//TODO: provide tests
func Init(file string) *Config {
	config, err := initConfig(file)
	if err != nil {
		log.Printf("Config file '%s' cannot be create (%s) \n", file, err)
	}

	return config
}

//TODO: provide tests
func initConfig(file string) (*Config, error) {
	cf := configFile{}

	err := yaml.Unmarshal(readFile(file), &cf)
	if err != nil {
		return &Config{}, err
	}

	config := new(Config)
	config.serverList = initServers(cf.Server)
	config.fileReport = cf.ReportFile

	return config, nil
}

// init servers struct; server url + disable params to report
func initServers(servers map[string]reportConfig) []Server {
	var serverList []Server
	for url, params := range servers {
		s, err := createServer(url, params.get())
		if err != nil {
			log.Println(err)
			continue
		}
		serverList = append(serverList, s)
	}

	return serverList
}

// read yaml file with configuration
//TODO: provide tests
func readFile(file string) []byte {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Unable to read config file: %v \n", err)
	}

	return []byte(yamlFile)
}
