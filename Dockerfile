# Use official Golang image as a builder
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy the backend source code
COPY rapatln_backend/ .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Use a small alpine image for the final run
FROM alpine:latest

# Working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
