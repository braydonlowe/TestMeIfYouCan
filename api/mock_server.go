package main 
 
type Book struct { 
    ID     int    `json:"id"` 
    Title  string `json:"title"` 
    Author string `json:"author"` 
} 
 
func main() { 
    http.HandleFunc("/books", booksHandler)       // GET & POST 
    http.HandleFunc("/books/", bookDeleteHandler) // DELETE 
    http.Handle("/", http.FileServer(http.Dir("./ui"))) 
    log.Fatal(http.ListenAndServe(":5000", nil)) 
}