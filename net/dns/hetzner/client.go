package hetzner

import (
	"net/http"
	"time"
)

type Client struct {
	key string
	hc  *http.Client
}

var (
	BaseURL = "https://dns.hetzner.com/api/v1"
)

// Return a new Client with specified timeout.
func NewClientWithTimeout(key string, timeout time.Duration) *Client {

	return &Client{key: key, hc: &http.Client{Timeout: timeout}}
}

// Return a new Client without timeout.
func NewClient(key string) *Client {
	return NewClientWithTimeout(key, 0)
}
