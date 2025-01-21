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

func CreateTransaction(apiClient *client.APIClient, req models.TransactionRequestUI) (*models.TransactionResponse, error) {
	log.Println("[INFO] Starting CreateTransaction")

	request := models.RootChargeCreditCardRequest{
		CreateTransactionRequest: models.CreateTransactionRequest{
			MerchantAuthentication: models.MerchantAuthentication{
				Name:           apiClient.Config.APILoginID,
				TransactionKey: apiClient.Config.TransactionKey,
			},
			TransactionRequest: models.TransactionRequest{
				TransactionType: "authCaptureTransaction", // इसे पहले रखें
				Amount:          req.TransactionAmount,
				Payment: models.Payment{
					CreditCard: models.CreditCard{
						CardNumber:     req.AccountNumber,
						ExpirationDate: req.ExpDate,
						CVV:            req.CVV,
					},
				},
				BillTo: models.BillTo{
					FirstName: req.BillingAddress.FirstName,
					LastName:  req.BillingAddress.LastName,
					Company:   req.BillingAddress.Company,
					Address:   req.BillingAddress.Street,
					City:      req.BillingAddress.City,
					State:     req.BillingAddress.State,
					Zip:       req.BillingAddress.PostalCode,
					Country:   req.BillingAddress.Country,
				},
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

	// // Check transaction response
	// if response.Messages.ResultCode != "Ok" {
	// 	if len(response.Messages.Message) > 0 {
	// 		log.Printf("[ERROR] Transaction failed with message: %s", response.Messages.Message[0].Text)
	// 		return nil, fmt.Errorf("transaction failed: %s", response.Messages.Message[0].Text)
	// 	}
	// 	log.Printf("[ERROR] Transaction failed with unknown error")
	// 	return nil, fmt.Errorf("transaction failed with unknown error")
	// }

	// log.Println("[INFO] Transaction completed successfully")
	return &response.TransactionResponse, nil
}
