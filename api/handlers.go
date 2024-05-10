package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/JoaoGumiero/Crud_Backend/ticket"
)

// Create a global map data structure to store tickets in memory and a mutex for concurrency read/write issues.
var (
	tickets = make(map[string]ticket.Ticket)
	mux     sync.RWMutex
)

// Function to retrieve all the tickets within the map.
func GetTickets(w http.ResponseWriter, r *http.Request) {
	mux.RLock()
	defer mux.RUnlock()
	var ticketList []ticket.Ticket
	for _, ticket := range tickets {
		ticketList = append(ticketList, ticket)
	}
	json.NewEncoder(w).Encode(ticketList)
}

func AddTicket(w http.ResponseWriter, r *http.Request) {
	var ticket ticket.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mux.Lock()
	tickets[ticket.ID] = ticket
	mux.Unlock()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ticket registered successfully"))
	json.NewEncoder(w).Encode(ticket)
}

func GetTicketById(w http.ResponseWriter, r *http.Request, id string) {
	// You can pass the id by the URL in the routes file or here with the following: id := r.URL.Path[len("/tickets/"):]
	// Lock and unlock reader concurrency [I've notes on Obsidian related to this]
	mux.RLock()
	defer mux.RUnlock()
	// Check if the ticket exist
	if _, ok := tickets[id]; !ok {
		http.Error(w, "Not Found", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(tickets[id])
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket ticket.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mux.Lock()
	defer mux.Unlock()
	if _, ok := tickets[ticket.ID]; !ok {
		http.Error(w, "Not Found", http.StatusMethodNotAllowed)
		return
	}
	tickets[ticket.ID] = ticket
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ticket updated successfully"))
	json.NewEncoder(w).Encode(ticket)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tickets/"):]
	mux.Lock()
	defer mux.Unlock()
	if _, ok := tickets[id]; !ok {
		http.Error(w, "Not Found", http.StatusMethodNotAllowed)
		return
	}
	delete(tickets, id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ticket deleted successfully"))
}
