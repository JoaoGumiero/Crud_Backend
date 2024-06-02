package api

import (
	"net/http"

	"github.com/JoaoGumiero/Crud_Backend/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Func to upload the routes into the new server, I've to make all the /tickets endpoints separated even with the Id, bcs the
// standard router don't support Id handling by itself.
func UploadRoutes(Dbpool *pgxpool.Pool) *http.ServeMux {
	ticketDao := postgres.NewTicketDAO(Dbpool)
	mux := http.NewServeMux()
	mux.HandleFunc("/tickets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			GetTickets(*ticketDao)(w, r)
		case "POST":
			AddTicket(*ticketDao)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tickets/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			UpdateTicket(*ticketDao)(w, r)
		case "DELETE":
			DeleteTicket(*ticketDao)(w, r)
		case "GET":
			// I could place the following code within the next route/switch case, but i preferred to place it here for learning purposes.
			GetTicketById(*ticketDao)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
