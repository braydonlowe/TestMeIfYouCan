#I'm not sure if I'm going to keep this because of the mingw32 that my computer is doing. I think I can do better with a .sh script.

.PHONY: test-api test-ui test-all

#Currently making me do mingw32-make as the command

#Go API tests
test-api:
	@printf "\n\n======= API Test Results =======\n"
	go test ./_tests/api/api_test.go

#TS tests
test-ui:
	@printf "\n\n======= UI Test Results =======\n"
	@echo "UI tests not implemented"

test-all:
	mingw32-make test-api test-ui
