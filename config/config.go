package config

type Config struct {
	APILoginID     string
	TransactionKey string
	BaseURL        string
}

func LoadConfig(isSandbox bool) Config {
	baseURL := "https://apitest.authorize.net/xml/v1/request.api"
	if !isSandbox {
		baseURL = "https://api.authorize.net/xml/v1/request.api"
	}

	return Config{
		APILoginID:     "9Z47qnW8",
		TransactionKey: "6uh7srY2r2G5K88r",
		BaseURL:        baseURL,
	}
}
