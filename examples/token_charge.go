package main

import (
	"authorize-net/client"
	"authorize-net/config"
	"authorize-net/models"
	"authorize-net/services"
	"log"
)

func main() {
	cfg := config.LoadConfig(true) // Pass `false` for production

	// Create API client
	apiClient := client.NewClient(cfg)

	// Build the request object
	req := models.CreateTokenTransactionRequest{
		TokenizeCreateTransactionRequest: models.TokenizeCreateTransactionRequest{
			MerchantAuthentication: models.MerchantAuthentication{
				Name:           "9Z47qnW8",
				TransactionKey: "6uh7srY2r2G5K88r",
			},
			TokenizeTransactionRequest: models.TokenizeTransactionRequest{
				TransactionType: "authCaptureTransaction",
				Amount:          "5.00",
				Profile: models.Profile{
					CustomerProfileID: "522074263",
					PaymentProfile: models.PaymentProfile{
						PaymentProfileID: "533955097",
					},
				},
			},
		},
	}

	// Call the service
	response, err := services.PerformTransactionViaToken(apiClient, req)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}

	log.Printf("Transaction successful! Response: %+v", response)
}
