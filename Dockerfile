FROM golang:1.24-alpine

WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire codebase
COPY . ./

# Set the working directory to the location of the tests
WORKDIR /app/_tests

# Run the Go test command
CMD ["go", "test", "./api/api_test.go"]
