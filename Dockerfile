# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

# Set up working directory
WORKDIR /app

# Install git (needed for go mod tidy in some cases)
RUN apk add --no-cache git

# Copy go.mod and go.sum, download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy everything else
COPY . ./

# Ensure the test script is copied correctly and list files
RUN ls -l /app

# Ensure tests.sh is executable
RUN chmod +x tests.sh

# Default command: run tests.sh
CMD ["./tests.sh"]
