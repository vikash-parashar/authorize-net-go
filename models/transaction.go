package models

type TransactionRequestUI struct {
	TransactionAmount string `json:"transaction_amount,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	ExpDate           string `json:"exp_date,omitempty"`
	CVV               string `json:"cvv,omitempty"`
	BillingAddress    struct {
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		Company    string `json:"company,omitempty"`
		Street     string `json:"street,omitempty"`
		City       string `json:"city,omitempty"`
		State      string `json:"state,omitempty"`
		PostalCode string `json:"postal_code,omitempty"`
		Country    string `json:"country,omitempty"`
	} `json:"billing_address,omitempty"`
	TransactionID string `json:"trans_id,omitempty"`
}

type RootChargeCreditCardRequest struct {
	CreateTransactionRequest CreateTransactionRequest `json:"createTransactionRequest,omitempty"`
}

type CreateTransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication,omitempty"`
	TransactionRequest     TransactionRequest     `json:"transactionRequest,omitempty"`
}

type MerchantAuthentication struct {
	Name           string `json:"name,omitempty"`
	TransactionKey string `json:"transactionKey,omitempty"`
}
type TransactionRequest struct {
	TransactionType string  `json:"transactionType,omitempty"`
	Amount          string  `json:"amount,omitempty"`
	Payment         Payment `json:"payment,omitempty"`
	BillTo          BillTo  `json:"billTo,omitempty"`
	// ShipTo          BillTo  `json:"shipTo,omitempty"`
	TransactionIdForVoidOrRefund string `json:"refTransId,omitempty"`
}

type Payment struct {
	CreditCard CreditCard `json:"creditCard,omitempty"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CVV            string `json:"cardCode,omitempty"`
}

type BillTo struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Company   string `json:"company,omitempty"`
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}

type RootChargeCreditCardResponse struct {
	TransactionResponse TransactionResponse `json:"transactionResponse,omitempty"`
	Messages            struct {
		ResultCode string    `json:"resultCode,omitempty"`
		Message    []Message `json:"message,omitempty"`
	}
}

type TransactionResponse struct {
	ResponseCode                           string    `json:"responseCode,omitempty"`
	AuthCode                               string    `json:"authCode,omitempty"`
	AvsResultCode                          string    `json:"avsResultCode,omitempty"`
	CvvResultCode                          string    `json:"cvvResultCode,omitempty"`
	CavvResultCode                         string    `json:"cavvResultCode,omitempty"`
	TransactionId                          string    `json:"transId,omitempty"`
	RefTransId                             string    `json:"refTransID,omitempty"`
	TransHash                              string    `json:"transHash,omitempty"`
	TestRequest                            string    `json:"testRequest,omitempty"`
	AccountNumber                          string    `json:"accountNumber,omitempty"`
	AccountType                            string    `json:"accountType,omitempty"`
	Messages                               []Message `json:"messages,omitempty"`
	Errors                                 []Error   `json:"errors,omitempty"`
	TransHashSha2                          string    `json:"transHashSha2,omitempty"`
	SupplementalDataQualificationIndicator int       `json:"SupplementalDataQualificationIndicator,omitempty"`
	NetworkTransId                         string    `json:"networkTransId,omitempty"`
}

type Message struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Text        string `json:"text,omitempty"`
}

type TransactionResponseUI struct {
	TransactionID       string `json:"transaction_id,omitempty"`
	ResponseCode        string `json:"response_code"`
	MessageDescription  string `json:"message_description"`
	MessageText         string `json:"message_text,omitempty"`
	MessageResponseCode string `json:"message_response_code"`
	Error               string `json:"errors,omitempty"`
}

type Error struct {
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorText string `json:"errorText,omitempty"`
}

type VoidTransactionRequest struct {
	CreateTransactionRequest struct {
		MerchantAuthentication struct {
			Name           string `json:"name,omitempty"`
			TransactionKey string `json:"transactionKey,omitempty"`
		} `json:"merchantAuthentication,omitempty"`
		TransactionRequest struct {
			TransactionType              string `json:"transactionType,omitempty"`
			TransactionIdForVoidOrRefund string `json:"refTransId,omitempty"`
		} `json:"transactionRequest,omitempty"`
	} `json:"createTransactionRequest,omitempty"`
}

type RefundTransactionRequest struct {
	CreateTransactionRequest struct {
		MerchantAuthentication struct {
			Name           string `json:"name,omitempty"`
			TransactionKey string `json:"transactionKey,omitempty"`
		} `json:"merchantAuthentication,omitempty"`
		TransactionRequest struct {
			TransactionType string `json:"transactionType,omitempty"`
			Amount          string `json:"amount,omitempty"`
			Payment         struct {
				CreditCard struct {
					CardNumber     string `json:"cardNumber,omitempty"`
					ExpirationDate string `json:"expirationDate,omitempty"`
					// CVV            string `json:"cardCode,omitempty"`
				} `json:"creditCard,omitempty"`
			} `json:"payment,omitempty"`
			TransactionIdForVoidOrRefund string `json:"refTransId,omitempty"`
		} `json:"transactionRequest,omitempty"`
	} `json:"createTransactionRequest,omitempty"`
}