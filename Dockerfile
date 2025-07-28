# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Run tests (only under ./api/)
CMD ["go", "test", "-v", "./api"]
