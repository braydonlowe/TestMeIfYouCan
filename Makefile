.PHONY: test-api test-ui test-all

#Currently making me do mingw32-make as the command

#Go API tests
test-api:
	@echo " \ \ =======API Test Results======="
	go test ./_tests/api/api_test.go

#TS tests
test-ui:
	@echo "UI tests not implemented"

test-all:
	mingw32-make test-api test-ui
