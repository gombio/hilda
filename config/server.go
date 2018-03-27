package config

import (
	"net/url"
)

type Server interface {
	GetUrl() string
	GetFlags() []string
}

type server struct {
	url           string
	disableParams []string
}

func createServer(address string, params []string) (server, error) {
	if _, err := url.ParseRequestURI(address); err != nil {
		return server{}, err
	}
	return server{
		url:           address,
		disableParams: params,
	}, nil
}

func (s server) GetUrl() string {
	return s.url
}

func (s server) GetFlags() []string {
	return s.disableParams
}
