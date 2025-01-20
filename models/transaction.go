package models

type TransactionRequest struct {
	MerchantAuthentication struct {
		Name           string `json:"name"`
		TransactionKey string `json:"transactionKey"`
	} `json:"merchantAuthentication"`
	TransactionType string `json:"transactionType"`
	Amount          string `json:"amount"`
	Payment         struct {
		CreditCard struct {
			CardNumber     string `json:"cardNumber"`
			ExpirationDate string `json:"expirationDate"`
			CVV            string `json:"cardCode"`
		} `json:"creditCard"`
	} `json:"payment"`
}

type TransactionResponse struct {
	TransactionID string `json:"transactionId"`
	ResponseCode  string `json:"responseCode"`
	Message       string `json:"message"`
}
