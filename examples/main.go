package main

// import (
// 	"authorize-net/client"
// 	"authorize-net/config"
// 	"authorize-net/models"
// 	"authorize-net/services"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	cfg := config.LoadConfig(true) // Pass `false` for production

// 	// Create API client
// 	apiClient := client.NewClient(cfg)

// 	// response, err := services.CreateTransaction(apiClient, models.TransactionRequestUI{
// 	// 	TransactionAmount: "5.00",
// 	// 	AccountNumber:     "4111111111111111",
// 	// 	ExpDate:           "2025-12",
// 	// 	CVV:               "999",
// 	// 	BillingAddress: struct {
// 	// 		FirstName  string `json:"first_name,omitempty"`
// 	// 		LastName   string `json:"last_name,omitempty"`
// 	// 		Company    string `json:"company,omitempty"`
// 	// 		Street     string `json:"street,omitempty"`
// 	// 		City       string `json:"city,omitempty"`
// 	// 		State      string `json:"state,omitempty"`
// 	// 		PostalCode string `json:"postal_code,omitempty"`
// 	// 		Country    string `json:"country,omitempty"`
// 	// 	}{
// 	// 		FirstName:  "John",
// 	// 		LastName:   "Doe",
// 	// 		Company:    "Authorize.Net",
// 	// 		Street:     "123 Main St",
// 	// 		City:       "Bellevue",
// 	// 		State:      "WA",
// 	// 		PostalCode: "98004",
// 	// 		Country:    "USA",
// 	// 	},
// 	// })

// 	// for refund only
// 	response, err := services.RefundTransaction(apiClient, models.TransactionRequestUI{
// 		TransactionAmount: "5.00",
// 		AccountNumber:     "4111111111111111",
// 		ExpDate:           "2025-12",
// 		CVV:               "999",
// 		TransactionID:     "80034779376",
// 	})

// 	if err != nil {
// 		log.Fatalf("Transaction failed: %v", err)
// 	}
// 	// log.Printf("Transaction successful! ID: %s", response)

// 	var res = models.TransactionResponseUI{
// 		TransactionID: response.TransactionId,
// 		ResponseCode:  response.ResponseCode,
// 	}

// 	if len(response.Messages) > 0 {
// 		res.MessageDescription = response.Messages[0].Description
// 		res.MessageText = response.Messages[0].Text
// 		res.MessageResponseCode = response.Messages[0].Code
// 	}

// 	if len(response.Errors) > 0 {
// 		res.Error = response.Errors[0].ErrorText
// 	}

// 	// Print the response
// 	fmt.Println("-------------------------------")
// 	fmt.Printf("Response: %+v\n", res)
// 	fmt.Println("-------------------------------")

// 	fmt.Println("-------------------------------")
// 	fmt.Println("response -----------: ", res)
// 	fmt.Println("-------------------------------")

// 	//TODO: for cancel a trxn
// 	// log.Println("processing refund ")
// 	// response, err := services.VoidTransaction(apiClient, "80034764695")

// 	// if err != nil {
// 	// 	log.Fatalf("Transaction failed: %v", err)
// 	// }
// 	// // log.Printf("Transaction successful! ID: %s", response)

// 	// var res = models.TransactionResponseUI{
// 	// 	TransactionID: response.TransactionId,
// 	// 	ResponseCode:  response.ResponseCode,
// 	// }

// 	// if len(response.Messages) > 0 {
// 	// 	res.MessageDescription = response.Messages[0].Description
// 	// 	res.MessageText = response.Messages[0].Text
// 	// 	res.MessageResponseCode = response.Messages[0].Code
// 	// }

// 	// if len(response.Errors) > 0 {
// 	// 	res.Error = response.Errors[0].ErrorText
// 	// }

// 	res2, err := json.Marshal(res)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	fmt.Println(string(res2))

// }
