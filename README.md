## Simple Customer Order Service in Golang

This is a simple go order service making use of gin framework. The below features are the ones i managed to get up and running

#### Features
- REST API to create customer, get customer details and manage customer orders
- Working with gorm.
- Docker and docker-compose for the dependencies
- CI/CD pipeline for automated testing and deployment

#### Prerequisites
Go (version 1.22+)
Git
Docker

#### Setup

Clone this repository: 
```
git clone https://github.com/kipngeno-isaac/go-order-service.git
```

Configure environment variables:
```
cp .env.example .env
```

Run docker-compose: 
```
docker-compose up
```
Build app image:
```
docker build -f Dockerfile -t go-order-service:latest .
```

#### Running the Service
Start the service:
``` 
docker run -d -p 3000:8080 go-order-service:latest
```

Access the API using your preferred HTTP client.

#### API Endpoints
**Create customer**
Sample request
```
curl --location 'http://localhost:3000/customers' \
--header 'Content-Type: application/json' \
--data '{
    "Username":"john",
    "Phone":"0723116789"
}'
```
sample response
```
{
    "customer": {
        "id": 1,
        "username": "john",
        "phone": "0723116789"
    }
}
```

**Create order**
sample request
```
curl --location 'http://localhost:3000/orders/1' \
--header 'Content-Type: application/json' \
--data '{
    "Item": "watch",
    "Amount":650
}'
```
Sample response
```
{
    "Order": {
        "id": 1,
        "item": "watch",
        "amount": 650,
        "customer_id": 1
    }
}
```
#### Missing Features
- Implement authentication and authorization via OpenID Connect.
- SMS alerting on successful order creation
- Data validation
- Testing and testing coverage implementation