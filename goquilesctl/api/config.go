package api

import (
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/latticework/proto-goquiles-server-todo/goquilesctl/command"
)

// @kenbrubaker: Goal here is to merge service and client.

var (
	errRedirect = errors.New("redirect")
)

// @kenbrubaker: This content is in client.go in vault.
type Config struct {
	// Client fields.
	Address string

	HttpClient *http.Client

	// Service fields.
	ServiceVersion string
}

func DefaultConfig() *Config {
	config := Config{
		Address:    "http://127.0.0.1:8700",
		HttpClient: &http.Client{},
	}

	if address := os.Getenv(command.EnvQuilesServerAddress); address != "" {
		config.Address = address
	}

	return config
}

func NewClient(c *Config) (*Client, error) {
	url, err := url.Parse(c.Address)
	if err != nil {
		return nil, err
	}

	if c.HttpClient == nil {
		c.HttpClient = http.DefaultClient
	}

	//	// Make a copy of the HTTP client so we can configure it without
	//	// affecting the original
	//	//
	//	// If no cookie jar is set on the client, we set a default empty
	//	// cookie jar.
	//	if c.HttpClient.Jar == nil {
	//		jar, err := cookiejar.New(&cookiejar.Options{})
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		c.HttpClient.Jar = jar
	//	}

	// Ensure redirects are not automatically followed
	c.HttpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errRedirect
	}

	client := &Client{
		address: url,
		config:  c,
	}

	//	if token := os.Getenv("VAULT_TOKEN"); token != "" {
	//		client.SetToken(token)
	//	}

	return client, nil
}
