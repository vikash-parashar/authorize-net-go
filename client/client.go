package client

import (
	"authorize-net/config"
	"bytes"
	"log"
	"net/http"
	"time"
)

type APIClient struct {
	Client  *http.Client
	Config  config.Config
	BaseURL string
}

func NewClient(cfg config.Config) *APIClient {
	return &APIClient{
		Client:  &http.Client{Timeout: 10 * time.Second},
		Config:  cfg,
		BaseURL: cfg.BaseURL,
	}
}

func (apiClient *APIClient) Post(endpoint string, payload []byte) (*http.Response, error) {
	url := apiClient.BaseURL + endpoint
	log.Printf("[DEBUG] Making POST request to URL: %s", url)
	log.Printf("[DEBUG] Payload: %s", string(payload))

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		log.Printf("[ERROR] Failed to create request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := apiClient.Client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to send request: %v", err)
		return nil, err
	}

	log.Printf("[DEBUG] Response Status: %d", resp.StatusCode)
	return resp, nil
}
