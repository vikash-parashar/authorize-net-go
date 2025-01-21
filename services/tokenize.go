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

func TokenizeCreditCard(apiClient *client.APIClient, req models.TokenizeCreditCardRequest) (*models.TokenizeCreditCardResponse, error) {
	log.Println("[INFO] Starting Tokenize Credit Card")

	// Serialize request to JSON
	payload, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Printf("[ERROR] Failed to marshal request payload: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] Serialized payload: %s", string(payload))

	// Send API request
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
	var response models.TokenizeCreditCardResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("[ERROR] Failed to decode API response: %v", err)
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	log.Printf("[DEBUG] Decoded Response: %+v", response)

	// Check transaction response
	if response.Messages.ResultCode != "Ok" {
		if len(response.Messages.Message) > 0 {
			log.Printf("[ERROR] Tokenization failed with message: %s", response.Messages.Message[0].Text)
			return nil, fmt.Errorf("tokenization failed: %s", response.Messages.Message[0].Text)
		}
		log.Printf("[ERROR] Tokenization failed with unknown error")
		return nil, fmt.Errorf("tokenization failed with unknown error")
	}

	log.Println("[INFO] Tokenization completed successfully")
	return &response, nil
}

func PerformTransactionViaToken(apiClient *client.APIClient, req models.CreateTokenTransactionRequest) (*models.TokenTransactionResponse, error) {
	log.Println("[INFO] Starting Perform Transaction")

	// Serialize request to JSON
	payload, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Printf("[ERROR] Failed to marshal request payload: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] Serialized payload: %s", string(payload))

	// Send API request
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
	var response models.TokenTransactionResponse
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
	return &response, nil
}
