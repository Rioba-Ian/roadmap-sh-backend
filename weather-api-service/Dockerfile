# Use the official Golang image as a base
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go modules
COPY weather-api-service/go.mod weather-api-service/go.sum ./

# Download go modules
RUN go mod download

# Copy the source code into the container
COPY . ./app


# Build the Go application
RUN go build -o main .

# Use a minimal image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
