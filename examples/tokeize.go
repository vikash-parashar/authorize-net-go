package main

// import (
// 	"authorize-net/client"
// 	"authorize-net/config"
// 	"authorize-net/models"
// 	"authorize-net/services"
// 	"log"
// )

// func main() {
// 	cfg := config.LoadConfig(true) // Pass `false` for production

// 	// Create API client
// 	apiClient := client.NewClient(cfg)

// 	// Build the request object
// 	req := models.TokenizeCreditCardRequest{
// 		CreateCustomerProfileRequest: struct {
// 			MerchantAuthentication struct {
// 				Name           string `json:"name,omitempty"`
// 				TransactionKey string `json:"transactionKey,omitempty"`
// 			} `json:"merchantAuthentication,omitempty"`
// 			Profile struct {
// 				MerchantCustomerID string `json:"merchantCustomerId,omitempty"`
// 				Description        string `json:"description,omitempty"`
// 				Email              string `json:"email,omitempty"`
// 				PaymentProfiles    struct {
// 					CustomerType string `json:"customerType,omitempty"`
// 					Payment      struct {
// 						CreditCard struct {
// 							CardNumber     string `json:"cardNumber,omitempty"`
// 							ExpirationDate string `json:"expirationDate,omitempty"`
// 							CardCode       string `json:"cardCode,omitempty"`
// 						} `json:"creditCard,omitempty"`
// 					} `json:"payment,omitempty"`
// 				} `json:"paymentProfiles,omitempty"`
// 			} `json:"profile,omitempty"`
// 			ValidationMode string `json:"validationMode,omitempty"`
// 		}{
// 			MerchantAuthentication: struct {
// 				Name           string `json:"name,omitempty"`
// 				TransactionKey string `json:"transactionKey,omitempty"`
// 			}{
// 				Name:           "9Z47qnW8",
// 				TransactionKey: "6uh7srY2r2G5K88r",
// 			},
// 			Profile: struct {
// 				MerchantCustomerID string `json:"merchantCustomerId,omitempty"`
// 				Description        string `json:"description,omitempty"`
// 				Email              string `json:"email,omitempty"`
// 				PaymentProfiles    struct {
// 					CustomerType string `json:"customerType,omitempty"`
// 					Payment      struct {
// 						CreditCard struct {
// 							CardNumber     string `json:"cardNumber,omitempty"`
// 							ExpirationDate string `json:"expirationDate,omitempty"`
// 							CardCode       string `json:"cardCode,omitempty"`
// 						} `json:"creditCard,omitempty"`
// 					} `json:"payment,omitempty"`
// 				} `json:"paymentProfiles,omitempty"`
// 			}{
// 				MerchantCustomerID: "CUST002",
// 				Description:        "Vikash Parashar",
// 				Email:              "customercare@example.com",
// 				PaymentProfiles: struct {
// 					CustomerType string `json:"customerType,omitempty"`
// 					Payment      struct {
// 						CreditCard struct {
// 							CardNumber     string `json:"cardNumber,omitempty"`
// 							ExpirationDate string `json:"expirationDate,omitempty"`
// 							CardCode       string `json:"cardCode,omitempty"`
// 						} `json:"creditCard,omitempty"`
// 					} `json:"payment,omitempty"`
// 				}{
// 					CustomerType: "individual",
// 					Payment: struct {
// 						CreditCard struct {
// 							CardNumber     string `json:"cardNumber,omitempty"`
// 							ExpirationDate string `json:"expirationDate,omitempty"`
// 							CardCode       string `json:"cardCode,omitempty"`
// 						} `json:"creditCard,omitempty"`
// 					}{
// 						CreditCard: struct {
// 							CardNumber     string `json:"cardNumber,omitempty"`
// 							ExpirationDate string `json:"expirationDate,omitempty"`
// 							CardCode       string `json:"cardCode,omitempty"`
// 						}{
// 							CardNumber:     "4007000000027",
// 							ExpirationDate: "2025-12",
// 							CardCode:       "123",
// 						},
// 					},
// 				},
// 			},
// 			ValidationMode: "testMode",
// 		},
// 	}

// 	// Call the service
// 	response, err := services.TokenizeCreditCard(apiClient, req)
// 	if err != nil {
// 		log.Fatalf("Tokenization failed: %v", err)
// 	}

// 	log.Printf("Tokenization successful! Response: %+v", response)
// }

// // CustomerProfileID:522074263 CustomerPaymentProfileIDList:[533955097]
