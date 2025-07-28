package api_test

import (
	//"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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


func TestGetBooks_AfterMultipleCreates(t *testing.T) {
	api.ResetBooks()

	booksToCreate := []string{
		`{"title": "Book One", "author": "Author A"}`,
		`{"title": "Book Two", "author": "Author B"}`,
	}

	for _, b := range booksToCreate {
		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		api.BooksHandler(w, req)
		if w.Result().StatusCode != http.StatusCreated {
			t.Fatalf("failed to create book: %s", b)
		}
	}

	// Check they were stored
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	api.BooksHandler(w, req)

	var books []api.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Fatalf("failed to decode books: %v", err)
	}
	if len(books) != 2 {
		t.Fatalf("expected 2 books, got %d", len(books))
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

func TestCreateBook_MissingAuthor(t *testing.T) {
	api.ResetBooks()
	payload := `{"title": "Missing Author"}`
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

//DELETE /books/<ID>
func TestDeleteBook_InvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/books/abc", nil)
	w := httptest.NewRecorder()

	api.BookDeleteHandler(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", w.Result().StatusCode)
	}
}

func TestDeleteBook_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/books/999", nil)
	w := httptest.NewRecorder()

	api.BookDeleteHandler(w, req)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404 Not Found, got %d", w.Result().StatusCode)
	}
}

func TestDeleteBook_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/books/1", nil) // Should be DELETE
	w := httptest.NewRecorder()
	api.BookDeleteHandler(w, req)

	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method Not Allowed, got %d", w.Result().StatusCode)
	}
}



func TestDeleteBook_Success(t *testing.T) {
	// Create book
	payload := `{"title": "Temp Book", "author": "Temp Author"}`
	createReq := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(payload))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()

	api.BooksHandler(createW, createReq)

	var createdBook api.Book
	json.NewDecoder(createW.Body).Decode(&createdBook)

	// Delete it
	deleteReq := httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(createdBook.ID), nil)
	deleteW := httptest.NewRecorder()
	api.BookDeleteHandler(deleteW, deleteReq)

	if deleteW.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", deleteW.Result().StatusCode)
	}
}


//BooksHandler test

func TestBooksHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/books", nil) // Unsupported method
	w := httptest.NewRecorder()
	api.BooksHandler(w, req)

	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 Method Not Allowed, got %d", w.Result().StatusCode)
	}
}

