package services

// import (
// 	"authorize-net/client"
// 	"authorize-net/models"
// 	"encoding/json"
// 	"fmt"
// )

// func VoidTransaction(apiClient *client.APIClient, transactionID string) (*models.TransactionResponse, error) {
// 	request := models.VoidRequest{
// 		MerchantAuthentication: models.MerchantAuthentication{
// 			Name:           apiClient.Config.APILoginID,
// 			TransactionKey: apiClient.Config.TransactionKey,
// 		},
// 		TransactionType: "voidTransaction",
// 		RefTransId:      transactionID,
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
// 		return nil, fmt.Errorf("void failed: %s", response.Message)
// 	}

// 	return &response, nil
// }
