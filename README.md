## Simple Customer Order Service in Golang

This repository contains a Go service for managing customers and orders with OpenID Connect authentication and SMS notifications.

#### Features
REST API for adding, updating, and retrieving customers and orders
OpenID Connect authentication for secure access
SMS notifications to customers on new orders using Africa's Talking API
Unit tests with coverage reporting
CI/CD pipeline for automated testing and deployment

#### Prerequisites
Go (version 1.18+)
Git
A cloud platform like Heroku, AWS Lambda, or Google Cloud Run
An Africa's Talking account (for sandbox testing)

#### Setup

Clone this repository: `git clone https://github.com/your-username/customer-order-service.git`
Install dependencies: `go mod download`
Configure environment variables:
```
OIDC_PROVIDER_URL: URL of your OpenID Connect provider
CLIENT_ID: Your OIDC client ID
CLIENT_SECRET: Your OIDC client secret
AFRICASTALKING_USERNAME: Your Africa's Talking username
AFRICASTALKING_API_KEY: Your Africa's Talking API key
```
Build the service: `go build -o cmd/customer-order-service ./cmd/customer-order-service`

#### Running the Service
Start the service:` ./cmd/customer-order-service`
Access the API using your preferred HTTP client.

#### API Endpoints
```
/customers: List, create, update, and delete customers
/orders: List, create, and update orders
```
#### Deployment

Follow the instructions for your chosen cloud platform to deploy the service.
Update environment variables with your platform-specific credentials.

#### Testing

Run unit tests: `go test -coverprofile=coverage.out -covermode=set && go tool cover -html=coverage.out`

#### Contributing

Pull requests are welcome! Please follow the contribution guidelines in the CONTRIBUTING.md file.