package services

import (
	"authorize-net/client"
	"authorize-net/models"
	"encoding/json"
	"fmt"
)

func CreateTransaction(apiClient *client.APIClient, request models.TransactionRequest) (*models.TransactionResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := apiClient.Post("/createTransactionRequest", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response models.TransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if response.ResponseCode != "1" {
		return nil, fmt.Errorf("transaction failed: %s", response.Message)
	}

	return &response, nil
}
