package api

import (
	"net/http"
)

func UploadRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			GetBooks(w, r)
		case "POST":
			AddBook(w, r)
		default:
			http.Error(w, "Method Not Allowed", 405)
		}
	})

	mux.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/books/"):]
		switch r.Method {
		case "PUT":
			UpdateBook(w, r)
		case "DELETE":
			DeleteBook(w, r)
		case "GET":
			// I could place the following code within the next route/switch case, but i preferred to place it here for learning purposes.
			GetBookById(w, r, id)
		default:
			http.Error(w, "Method Not Allowed", 405)
		}
	})
}
