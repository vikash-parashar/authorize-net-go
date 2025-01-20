# Authorize.Net Golang Module

Welcome to the **Authorize.Net Golang Module**, a lightweight and modular package for integrating Authorize.Net payment gateway into your Golang web applications. This package is designed to support a variety of payment operations, including recurring payments, terminal integrations, refunds, and cancellations.

## Repository
GitHub Repository: [authorize-net-go](https://github.com/vikash-parashar/authorize-net-go)

---

## Features

- Modular and reusable architecture for seamless integration.
- Supports basic and advanced Authorize.Net REST API functionalities.
- Includes utilities for transaction handling, recurring payments, and logging.
- Well-documented examples for ease of use.

---

## Directory Structure

```plaintext
|-- authorize-net/
    |-- config/                # Configuration files for API credentials
    |-- client/                # HTTP client for API requests
    |-- models/                # Data models for requests and responses
    |-- services/              # Core logic for payment operations
    |-- utils/                 # Helper functions and utilities
    |-- examples/              # Example implementations
    |-- go.mod                 # Go module file
    |-- go.sum                 # Dependency tracking
```

---

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/vikash-parashar/authorize-net-go.git
    ```

2. Navigate to the project directory:
    ```bash
    cd authorize-net-go
    ```

3. Install dependencies:
    ```bash
    go mod tidy
    ```

4. Import the module in your project.

---

## Configuration

Update the `config/config.go` file with your Authorize.Net credentials:

```go
return Config{
    APILoginID:    "YOUR_API_LOGIN_ID",
    TransactionKey: "YOUR_TRANSACTION_KEY",
    BaseURL:       "https://apitest.authorize.net/xml/v1/request.api",
}
```

---

## Usage

Here is an example of creating a transaction:

```go
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
```

---

## Contributions

Contributions are welcome! Feel free to fork this repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

---

## Author

- **Name**: Vikash Parashar
- **LinkedIn**: [Vikash Parashar](https://www.linkedin.com/in/vikash-parashar-3152471ba/)
- **Email**: [gowithvikash@gmail.com](mailto:gowithvikash@gmail.com)

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

## Acknowledgements

Special thanks to the Golang and Authorize.Net communities for their amazing documentation and support!

