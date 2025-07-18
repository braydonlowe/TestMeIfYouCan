package qa_api_tests

import "testing"


func AlwaysTrue() bool {
	return true
}

func TestAlwaysTrue(t *testing.T) {
	if !AlwaysTrue() {
		t.Error("AlwaysTrue() returned false, expected true")
	}
}