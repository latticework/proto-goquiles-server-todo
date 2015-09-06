package api

import (
	"net/url"
)

type Client struct {
	address *url.URL
	config  *Config
}
