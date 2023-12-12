package client

import (
	"devgym-http/src/service"
	"net/http"
	"time"
)

const prodUrl = "http://vascodagama.com.br"

type Client struct {
	url        string
	timeout    time.Duration
	httpClient *http.Client

	Deployment service.Service
}

func NewClient(options ...option) (*Client, error) {
	c := Client{
		url:        prodUrl,
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

	c.Deployment = service.NewService(c.httpClient, c.url)

	return &c, nil
}
