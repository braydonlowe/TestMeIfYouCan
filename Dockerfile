# syntax=docker/dockerfile:1

FROM golang:1.21-alpine

# Set up working directory
WORKDIR /app

# Install git (needed for go mod tidy in some cases)
RUN apk add --no-cache git

# Copy go.mod and go.sum, download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy everything else
COPY . .

# Run tests by default
CMD ["go", "test", "-v", "./..."]
