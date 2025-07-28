package qa_api_tests

import (
	//"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	//"strconv"
	//"strings"
	"testing"
	"github.com/braydonlowe/TestMeIfYouCan/api"
)

//Makes sure that my tests are running properly.
func AlwaysTrue() bool {
	return true
}

func TestAlwaysTrue(t *testing.T) {
	if !AlwaysTrue() {
		t.Error("AlwaysTrue() returned false, expected true")
	}
}

func TestGetBooks_InitiallyEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()

	api.BooksHandler(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", w.Result().StatusCode)
	}

	var books []api.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Fatalf("failed to decode reqpose: %v", err)
	}

	if len(books) != 0 {
		t.Errorf("expected empty book list, got %d", len(books))
	}
}