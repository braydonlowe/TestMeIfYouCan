package qa_api_tests

import (
	//"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	//"strconv"
	"strings"
	"testing"
	"github.com/braydonlowe/TestMeIfYouCan/api"
)


//GET /books
func TestGetBooks_InitiallyEmpty(t *testing.T) {
	api.ResetBooks() // Custom function
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()

	api.BooksHandler(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", w.Result().StatusCode)
	}

	var books []api.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	//Tests that this still works when the db is empty.
	if len(books) != 0 {
		t.Errorf("expected empty book list, got %d", len(books))
	}
}

//POST /books
func TestCreateBook_Valid(t *testing.T) {
	payload := `{
				"title": "Othello",
				"author": "William Shakespeare"
				}`

	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	api.BooksHandler(w, req)

	if w.Result().StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201 Created, got %d", w.Result().StatusCode)
	}

	//Take the book and make sure it's stored.
	var book api.Book
	if err := json.NewDecoder(w.Body).Decode(&book); err != nil {
		t.Fatalf("faled to decode book: %v", err)
	}

	if book.Title != "Othello" || book.Author != "William Shakespeare" {
		t.Errorf("unexpected book content: %+v", book)
	}
}


func TestCreateBook_missingTitle(t *testing.T) {
	payload := `{"author": "No Title"}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	api.BooksHandler(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", w.Result().StatusCode)
	}
}

func TestCreateBook_InvalidJSON(t *testing.T) {
	payload := `This is not a json at all`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	api.BooksHandler(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", w.Result().StatusCode)
	}
}