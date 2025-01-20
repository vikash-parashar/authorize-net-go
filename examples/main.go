package main

import (
	"authorize-net/client"
	"authorize-net/config"
	"authorize-net/models"
	"authorize-net/services"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()
	apiClient := client.NewClient(cfg.BaseURL)

	transactionRequest := models.TransactionRequest{
		TransactionType: "authCaptureTransaction",
		Amount:          "100.00",
	}
	transactionRequest.MerchantAuthentication.Name = cfg.APILoginID
	transactionRequest.MerchantAuthentication.TransactionKey = cfg.TransactionKey
	transactionRequest.Payment.CreditCard.CardNumber = "4111111111111111"
	transactionRequest.Payment.CreditCard.ExpirationDate = "2025-12"
	transactionRequest.Payment.CreditCard.CVV = "123"

	response, err := services.CreateTransaction(apiClient, transactionRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Transaction Successful! ID: %s\n", response.TransactionID)
}
