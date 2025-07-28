package api


import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

var (
    books       = []Book{}
    nextID      = 1
    booksMutex sync.Mutex
    //For concurrency I could use sync.RWMutex then lock or unlock read or write
)

//Since I defined this file, I've added this. I'll test at scale as well, but I want to test everything individually.
func ResetBooks() {
	booksMutex.Lock()
	defer booksMutex.Unlock()
	books = []Book{}
	nextID = 1
}

//Handlers:
func BooksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:    //method given in the challenge code
            getBooks(w, r)
        case http.MethodPost:   //Method given in the challenge code
            createBook(w, r)
        default:
            http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
    }
}


func BookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	booksMutex.Lock()
	defer booksMutex.Unlock()

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

//Method implementations
func getBooks(w http.ResponseWriter, r *http.Request) {
    booksMutex.Lock()
    defer booksMutex.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
    var b Book
    if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if strings.TrimSpace(b.Title) == "" || strings.TrimSpace(b.Author) == "" {
        http.Error(w, "Missing title or author", http.StatusBadRequest)
        return
    }

    booksMutex.Lock()
    defer booksMutex.Unlock()

    b.ID = nextID
    nextID++
    books = append(books, b)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(b)
}

// From here down is the origional code that was given for the code challenge. See ReadMe for details.
type Book struct { 
    ID     int    `json:"id"` 
    Title  string `json:"title"` 
    Author string `json:"author"` 
} 
 
func main() { 
    http.HandleFunc("/books", BooksHandler)       // GET & POST 
    http.HandleFunc("/books/", BookDeleteHandler) // DELETE 
    http.Handle("/", http.FileServer(http.Dir("./ui"))) 
    log.Fatal(http.ListenAndServe(":5000", nil)) 
}