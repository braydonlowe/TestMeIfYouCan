#!/bin/bash

set -e


case "$1" in
    #Go API tests
    api)
        echo ""
        echo ""
        echo "======= API Test Results ======="
	    go test ./_tests/api/api_test.go
    ;;

    #TS tests
    ui)
        echo ""
        echo ""
	    echo "======= UI Test Results ======="
	    echo "UI tests not implemented"
    ;;

    #All tests
    all|"")
	    #test-api test-ui
        #From here down we copy and paste the top two
        echo ""
        echo "======= API Test Results ======="
        time go test ./_tests/api/api_test.go
        echo ""
        echo ""
        echo "======= UI Test Results ======="
        echo "UI tests not implemented"
    ;;
    *)
        echo "Usage: ./test.sh [api|ui|all]"
        exit 1
    ;;
esac