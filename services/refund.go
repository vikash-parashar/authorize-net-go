package services

// import (
// 	"authorize-net/client"
// 	"authorize-net/models"
// 	"encoding/json"
// 	"fmt"
// )

// func RefundTransaction(apiClient *client.APIClient, transactionID, amount, lastFourCardDigits string) (*models.TransactionResponse, error) {
// 	request := models.RefundRequest{
// 		MerchantAuthentication: models.MerchantAuthentication{
// 			Name:           apiClient.Config.APILoginID,
// 			TransactionKey: apiClient.Config.TransactionKey,
// 		},
// 		TransactionType: "refundTransaction",
// 		Amount:          amount,
// 		Payment: models.Payment{
// 			CreditCard: models.CreditCard{
// 				CardNumber: lastFourCardDigits,
// 			},
// 		},
// 		RefTransId: transactionID,
// 	}

// 	payload, err := json.Marshal(request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := apiClient.Post("/createTransactionRequest", payload)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var response models.TransactionResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return nil, err
// 	}

// 	if response.ResponseCode != "1" {
// 		return nil, fmt.Errorf("refund failed: %s", response.Message)
// 	}

// 	return &response, nil
// }
