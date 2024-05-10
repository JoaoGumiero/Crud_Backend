package api

import (
	"net/http"
)

func UploadRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tickets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			GetTickets(w, r)
		case "POST":
			AddTicket(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tickets/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/tickets/"):]
		switch r.Method {
		case "PUT":
			UpdateTicket(w, r)
		case "DELETE":
			DeleteTicket(w, r)
		case "GET":
			// I could place the following code within the next route/switch case, but i preferred to place it here for learning purposes.
			GetTicketById(w, r, id)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
