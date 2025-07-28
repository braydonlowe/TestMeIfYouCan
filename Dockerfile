# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

# Set up working directory
WORKDIR /app

# Install git (needed for some go get packages)
RUN apk add --no-cache git

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Explicitly copy the tests.sh
COPY tests.sh ./

# Ensure test script is executable (optional, if you use run-tests.sh)
RUN chmod +x tests.sh

# Default command: run the test script (you can still override this in GitHub Actions or locally)
CMD ["./tests.sh"]
