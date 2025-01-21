package models

// MerchantAuthentication represents authentication details.
type MerchantAuthentication struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

// CreditCard represents the payment card details.
type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CVV            string `json:"cardCode,omitempty"`
}

// LineItem represents an individual line item in the order.
type LineItem struct {
	ItemID      string  `json:"itemId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

type BillTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

type ShipTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Company   string `json:"company"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

// TransactionRequestDetails represents the main details of the transaction request.
type TransactionRequestDetails struct {
	TransactionType string `json:"transactionType"`
	Amount          string `json:"amount"`
	Payment         struct {
		CreditCard CreditCard `json:"creditCard"`
	} `json:"payment"`
	Order struct {
		InvoiceNumber string `json:"invoiceNumber,omitempty"`
		Description   string `json:"description,omitempty"`
	} `json:"order,omitempty"`

	Shipping struct {
		Amount      float64 `json:"amount"`
		Name        string  `json:"name,omitempty"`
		Description string  `json:"description,omitempty"`
	} `json:"shipping,omitempty"`
	// Customer struct {
	// 	ID string `json:"id"`
	// } `json:"customer,omitempty"`
	BillTo BillTo `json:"billTo,omitempty"`
	ShipTo ShipTo `json:"shipTo,omitempty"`
}

// RootRequest represents the full JSON request.
// RootRequest represents the full JSON request.
type RootRequest struct {
	CreateTransactionRequest struct {
		// RefId                  string                    `json:"refId,omitempty"`
		MerchantAuthentication MerchantAuthentication    `json:"merchantAuthentication"`
		TransactionRequest     TransactionRequestDetails `json:"transactionRequest"`
	} `json:"createTransactionRequest"`
}

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
}

// Messages represent the overall status and messages of the response.
type Messages struct {
	ResultCode string `json:"resultCode"`
	Message    []struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"message"`
}

// TransactionResponse represents the main transaction details.
type TransactionResponse struct {
	ResponseCode   string `json:"responseCode,omitempty"`
	AuthCode       string `json:"authCode,omitempty"`
	AVSResultCode  string `json:"avsResultCode,omitempty"`
	CVVResultCode  string `json:"cvvResultCode,omitempty"`
	CAVVResultCode string `json:"cavvResultCode,omitempty"`
	TransID        string `json:"transId,omitempty"`
	RefTransID     string `json:"refTransID,omitempty"`
	TransHash      string `json:"transHash,omitempty"`
	TestRequest    string `json:"testRequest,omitempty"`
	AccountNumber  string `json:"accountNumber,omitempty"`
	AccountType    string `json:"accountType,omitempty"`
	Messages       []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"messages,omitempty"`
	UserFields []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"userFields,omitempty"`
	TransHashSha2                          string `json:"transHashSha2,omitempty"`
	SupplementalDataQualificationIndicator int    `json:"SupplementalDataQualificationIndicator,omitempty"`
	NetworkTransID                         string `json:"networkTransId,omitempty"`
}

// CreateTransactionResponse represents the full response for createTransaction.
type CreateTransactionResponse struct {
	RefID               string              `json:"refId,omitempty"`
	Messages            Messages            `json:"messages"`
	TransactionResponse TransactionResponse `json:"transactionResponse"`
}
