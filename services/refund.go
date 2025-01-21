package services

import (
	"authorize-net/client"
	"authorize-net/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func RefundTransaction(apiClient *client.APIClient, req models.TransactionRequestUI) (*models.TransactionResponse, error) {
	log.Println("[INFO] Starting Refund Transaction")

	// Build the request object
	request := models.RefundTransactionRequest{
		CreateTransactionRequest: struct {
			MerchantAuthentication struct {
				Name           string `json:"name,omitempty"`
				TransactionKey string `json:"transactionKey,omitempty"`
			} `json:"merchantAuthentication,omitempty"`
			TransactionRequest struct {
				TransactionType string `json:"transactionType,omitempty"`
				Amount          string `json:"amount,omitempty"`
				Payment         struct {
					CreditCard struct {
						CardNumber     string `json:"cardNumber,omitempty"`
						ExpirationDate string `json:"expirationDate,omitempty"`
					} `json:"creditCard,omitempty"`
				} `json:"payment,omitempty"`
				TransactionIdForVoidOrRefund string `json:"refTransId,omitempty"`
			} `json:"transactionRequest,omitempty"`
		}{
			MerchantAuthentication: struct {
				Name           string `json:"name,omitempty"`
				TransactionKey string `json:"transactionKey,omitempty"`
			}{
				Name:           apiClient.Config.APILoginID,
				TransactionKey: apiClient.Config.TransactionKey,
			},
			TransactionRequest: struct {
				TransactionType string `json:"transactionType,omitempty"`
				Amount          string `json:"amount,omitempty"`
				Payment         struct {
					CreditCard struct {
						CardNumber     string `json:"cardNumber,omitempty"`
						ExpirationDate string `json:"expirationDate,omitempty"`
					} `json:"creditCard,omitempty"`
				} `json:"payment,omitempty"`
				TransactionIdForVoidOrRefund string `json:"refTransId,omitempty"`
			}{
				TransactionType: "refundTransaction",
				Amount:          req.TransactionAmount,
				Payment: struct {
					CreditCard struct {
						CardNumber     string `json:"cardNumber,omitempty"`
						ExpirationDate string `json:"expirationDate,omitempty"`
					} `json:"creditCard,omitempty"`
				}{
					CreditCard: struct {
						CardNumber     string `json:"cardNumber,omitempty"`
						ExpirationDate string `json:"expirationDate,omitempty"`
					}{
						CardNumber:     req.AccountNumber,
						ExpirationDate: req.ExpDate,
					},
				},
				TransactionIdForVoidOrRefund: req.TransactionID,
			},
		},
	}

	// Serialize to JSON
	payload, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		log.Printf("[ERROR] Failed to marshal request payload: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] Serialized payload: %s", string(payload))

	// Send request
	resp, err := apiClient.Post(payload)
	if err != nil {
		log.Printf("[ERROR] API request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Log response status
	log.Printf("[DEBUG] Response Status: %d", resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read response body: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] Raw Response Body: %s", string(body))

	// Strip BOM if present
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	// Check Content-Type
	contentType := resp.Header.Get("Content-Type")
	log.Printf("[DEBUG] Content-Type: %s", contentType)
	if !strings.HasPrefix(contentType, "application/json") {
		log.Printf("[ERROR] Unexpected Content-Type: %s", contentType)
		return nil, fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	// Decode JSON response
	var response models.RootChargeCreditCardResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("[ERROR] Failed to decode API response: %v", err)
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	log.Printf("[DEBUG] Decoded Response: %+v", response)

	// Check transaction response
	if response.Messages.ResultCode != "Ok" {
		if len(response.Messages.Message) > 0 {
			log.Printf("[ERROR] Transaction failed with message: %s", response.Messages.Message[0].Text)
			return nil, fmt.Errorf("transaction failed: %s", response.Messages.Message[0].Text)
		}
		log.Printf("[ERROR] Transaction failed with unknown error")
		return nil, fmt.Errorf("transaction failed with unknown error")
	}

	log.Println("[INFO] Transaction completed successfully")
	return &response.TransactionResponse, nil
}
