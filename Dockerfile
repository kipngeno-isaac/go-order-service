# Set base image
FROM golang:1.22

# Set the base working directory for app 
WORKDIR /app

# Copy and download go modules
COPY go.* ./
RUN go mod download

# Copy the rest of the source code into working directory
COPY *.go ./

# Build image
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-order-service

# Set port to run container
EXPOSE 3000

# Run the container
CMD [ "/go-order-service" ]