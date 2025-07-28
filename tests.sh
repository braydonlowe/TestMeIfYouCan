#!/bin/bash

echo ""
echo ""
echo "======= API Test Results ======="
#go run api/mock_server.go
go test ./api/api_test.go