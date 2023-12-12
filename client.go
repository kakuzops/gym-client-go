package client

import (
	"net/http"
	"time"
)

type Client struct {
	url        string
	timeout    time.Duration
	httpClient *http.Client

	Deployment deployment.Service
}

func newClient(options ...option) (*Client, error) {
	c := Client{
		url:        "google.com.br",
		httpClient: &http.Client{},
	}

	for _, option := range options {
		if err := option(&c); err != nil {
			return nil, err
		}
	}

	if c.timeout != 0 {
		c.httpClient.Timeout = c.timeout
	}
}
