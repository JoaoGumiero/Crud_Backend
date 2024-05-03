package api

import (
	"encoding/json"
	"net/http"
	"sync"
)

// A book representation with an ID, Title and Author
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Create a global map data structure to store books in memory and a mutex for concurrency read/write issues.
var (
	books = make(map[string]Book)
	mux   sync.RWMutex
)

// Function to retrieve all the books within the map.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	mux.RLock()
	defer mux.RUnlock()
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	json.NewEncoder(w).Encode(bookList)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mux.Lock()
	books[book.ID] = book
	mux.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func GetBookById(w http.ResponseWriter, r *http.Request, id string) {
	// You can pass the id by the URL in the routes file or here with the following: id := r.URL.Path[len("/books/"):]
	// Lock and unlock reader concurrency [I've notes on Obsidian related to this]
	mux.RLock()
	defer mux.RUnlock()
	// Check if the book exist
	if _, ok := books[id]; !ok {
		http.Error(w, "Not Found", 404)
		return
	}
	json.NewEncoder(w).Encode(books[id])
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mux.Lock()
	defer mux.Unlock()
	if _, ok := books[book.ID]; !ok {
		http.Error(w, "Not Found", 404)
		return
	}
	books[book.ID] = book
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	mux.Lock()
	defer mux.Unlock()
	if _, ok := books[id]; !ok {
		http.Error(w, "Not Found", 404)
		return
	}
	delete(books, id)
}
