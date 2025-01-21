package services

// import (
// 	"authorize-net/client"
// 	"authorize-net/models"
// 	"encoding/json"
// 	"fmt"
// )

// func TokenizeCard(apiClient *client.APIClient, cardNumber, expirationDate, cvv string) (string, error) {
// 	request := models.TokenizeCardRequest{
// 		MerchantAuthentication: models.MerchantAuthentication{
// 			Name:           apiClient.Config.APILoginID,
// 			TransactionKey: apiClient.Config.TransactionKey,
// 		},
// 		Payment: models.Payment{
// 			CreditCard: models.CreditCard{
// 				CardNumber:     cardNumber,
// 				ExpirationDate: expirationDate,
// 				CVV:            cvv,
// 			},
// 		},
// 	}

// 	payload, err := json.Marshal(request)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := apiClient.Post("/createCustomerPaymentProfile", payload)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	var response models.TokenizeCardResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return "", err
// 	}

// 	if response.ResponseCode != "1" {
// 		return "", fmt.Errorf("failed to tokenize card: %s", response.Message)
// 	}

// 	return response.Token, nil
// }
