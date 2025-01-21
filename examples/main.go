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

	response, err := services.CreateTransaction(apiClient, models.TransactionRequestUI{
		TransactionAmount: "5.00",
		AccountNumber:     "4111111111111111",
		ExpDate:           "2025-12",
		CVV:               "999",
		BillingAddress: struct {
			FirstName  string `json:"first_name,omitempty"`
			LastName   string `json:"last_name,omitempty"`
			Company    string `json:"company,omitempty"`
			Street     string `json:"street,omitempty"`
			City       string `json:"city,omitempty"`
			State      string `json:"state,omitempty"`
			PostalCode string `json:"postal_code,omitempty"`
			Country    string `json:"country,omitempty"`
		}{
			FirstName:  "John",
			LastName:   "Doe",
			Company:    "Authorize.Net",
			Street:     "123 Main St",
			City:       "Bellevue",
			State:      "WA",
			PostalCode: "98004",
			Country:    "USA",
		},
	})
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}
	log.Printf("Transaction successful! ID: %s", response)

}
