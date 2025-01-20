package config

type Config struct {
	APILoginID     string
	TransactionKey string
	BaseURL        string
}

func LoadConfig() Config {
	return Config{
		APILoginID:     "YOUR_API_LOGIN_ID",
		TransactionKey: "YOUR_TRANSACTION_KEY",
		BaseURL:        "https://apitest.authorize.net/xml/v1/request.api",
	}
}
