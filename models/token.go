package models

type TokenizeCreditCardRequest struct {
	CreateCustomerProfileRequest struct {
		MerchantAuthentication struct {
			Name           string `json:"name,omitempty"`
			TransactionKey string `json:"transactionKey,omitempty"`
		} `json:"merchantAuthentication,omitempty"`
		Profile struct {
			MerchantCustomerID string `json:"merchantCustomerId,omitempty"`
			Description        string `json:"description,omitempty"`
			Email              string `json:"email,omitempty"`
			PaymentProfiles    struct {
				CustomerType string `json:"customerType,omitempty"`
				Payment      struct {
					CreditCard struct {
						CardNumber     string `json:"cardNumber,omitempty"`
						ExpirationDate string `json:"expirationDate,omitempty"`
						CardCode       string `json:"cardCode,omitempty"`
					} `json:"creditCard,omitempty"`
				} `json:"payment,omitempty"`
			} `json:"paymentProfiles,omitempty"`
		} `json:"profile,omitempty"`
		ValidationMode string `json:"validationMode,omitempty"`
	} `json:"createCustomerProfileRequest,omitempty"`
}

type TokenizeCreditCardResponse struct {
	CustomerProfileID             string   `json:"customerProfileId,omitempty"`
	CustomerPaymentProfileIDList  []string `json:"customerPaymentProfileIdList,omitempty"`
	CustomerShippingAddressIDList []string `json:"customerShippingAddressIdList,omitempty"`
	ValidationDirectResponseList  []string `json:"validationDirectResponseList,omitempty"`
	Messages                      struct {
		ResultCode string `json:"resultCode,omitempty"`
		Message    []struct {
			Code string `json:"code,omitempty"`
			Text string `json:"text,omitempty"`
		} `json:"message,omitempty"`
	} `json:"messages,omitempty"`
}

type CreateTokenTransactionRequest struct {
	TokenizeCreateTransactionRequest TokenizeCreateTransactionRequest `json:"createTransactionRequest,omitempty"`
}

type TokenizeCreateTransactionRequest struct {
	MerchantAuthentication     MerchantAuthentication     `json:"merchantAuthentication,omitempty"`
	TokenizeTransactionRequest TokenizeTransactionRequest `json:"transactionRequest,omitempty"`
}

type TokenizeTransactionRequest struct {
	TransactionType string  `json:"transactionType,omitempty"`
	Amount          string  `json:"amount,omitempty"`
	Profile         Profile `json:"profile,omitempty"`
}

type Profile struct {
	CustomerProfileID string         `json:"customerProfileId,omitempty"`
	PaymentProfile    PaymentProfile `json:"paymentProfile,omitempty"`
}

type PaymentProfile struct {
	PaymentProfileID string `json:"paymentProfileId,omitempty"`
}

type TokenTransactionResponse struct {
	TransactionResponse struct {
		ResponseCode  string `json:"responseCode"`
		AuthCode      string `json:"authCode"`
		AVSResultCode string `json:"avsResultCode"`
		CVVResultCode string `json:"cvvResultCode"`
		TransID       string `json:"transId"`
		Messages      []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"messages"`
	} `json:"transactionResponse"`
	Messages struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}
