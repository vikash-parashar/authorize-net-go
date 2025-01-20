package client

import (
	"bytes"
	"net/http"
	"time"
)

type APIClient struct {
	Client  *http.Client
	BaseURL string
}

func NewClient(baseURL string) *APIClient {
	return &APIClient{
		Client:  &http.Client{Timeout: 10 * time.Second},
		BaseURL: baseURL,
	}
}

func (c *APIClient) Post(endpoint string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
