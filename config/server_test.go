package config

import "testing"

func TestInitServer(t *testing.T) {
	url := "http://example.com"
	_, err := createServer(url, []string{})

	if err != nil {
		t.Fatalf("Cannot create server for '%s' url", url)
	}
}

func TestFailInitServer(t *testing.T) {
	url := "wrong-url"
	_, err := createServer(url, []string{})

	if err == nil {
		t.Fatalf("Server with url '%s' cannot exist but does", url)
	}
}

func TestUrlCorrectServer(t *testing.T) {
	url := "http://example.com/"
	server, _ := createServer(url, []string{})

	if server.GetURL() != "http://example.com/" {
		t.Fatalf("Server url should be proper but is equal to: '%s'", server.GetURL())
	}
}
