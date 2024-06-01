package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/JoaoGumiero/Crud_Backend/postgres"
	"github.com/JoaoGumiero/Crud_Backend/ticket"
)

// Create a global map data structure to store tickets in memory and a mutex for concurrency read/write issues.
var (
	tickets = make(map[string]ticket.Ticket)
	mux     sync.RWMutex
)

// Util for get id from path
func getIdFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("Invalid Path")
	}
	return strconv.Atoi(parts[2])
}

// Function to retrieve all the tickets within the map.
func GetTickets(t postgres.TicketDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Get all Tickets Handler")
		if r.Method != http.MethodGet {
			log.Fatalf("Method not allowed %d", http.StatusBadRequest)
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
		tickets, err := t.GetAllTicketsDAO(r.Context())
		if err != nil {
			log.Fatalf("Error retrieving all Tickets: %d", http.StatusBadRequest)
			http.Error(w, "Error retrieving all Tickets:", http.StatusBadRequest)
		}
		ticketJson, err := json.Marshal(tickets)
		if err != nil {
			log.Fatalf("Error marshling all tickets: %d", http.StatusBadRequest)
			http.Error(w, "Error marshling all tickets", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ticketJson)
		log.Print("Getting all Ticket process was successfull")
	}
}

func AddTicket(t postgres.TicketDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Add Tickets Handler")
		if r.Method != http.MethodPost {
			log.Fatalf("Method not allowed %d", http.StatusBadRequest)
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
		var ticket ticket.Ticket
		err := json.NewDecoder(r.Body).Decode(&ticket)
		if err != nil {
			log.Fatalf("Error decoding body: %d", http.StatusBadRequest)
			http.Error(w, "Error decoding body", http.StatusBadRequest)
		}
		createdTicket, err := t.CreateTicketDAO(r.Context(), ticket)
		if err != nil {
			log.Fatalf("Error creating the ticket: %d", http.StatusBadRequest)
			http.Error(w, "Error creating the ticket", http.StatusBadRequest)
		}
		ticketJson, err := json.Marshal(createdTicket)
		if err != nil {
			log.Fatalf("Error marshling tickets: %d", http.StatusBadRequest)
			http.Error(w, "Error marshling tickets", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ticketJson)
		log.Print("Ticket creating process was successfull")
	}
}

func GetTicketById(t postgres.TicketDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Get Tickets Handler by Id")
		if r.Method != http.MethodGet {
			log.Fatalf("Method not allowed %d", http.StatusBadRequest)
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
		id, err := getIdFromPath(r.URL.Path)
		if err != nil {
			log.Fatalf("Error getting id from Path: %d", http.StatusBadRequest)
			http.Error(w, "Error getting id from Path", http.StatusBadRequest)
		}
		ticket, err := t.GetTicketByIdDAO(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Fatalf("No ticket found with given Id: %d", http.StatusNotFound)
				http.Error(w, "No ticket found with given Id", http.StatusBadRequest)
			} else {
				log.Fatalf("No ticket found with given Id: %d", http.StatusNotFound)
				http.Error(w, "No ticket found with given Id", http.StatusBadRequest)
			}
		}
		ticketJson, err := json.Marshal(ticket)
		if err != nil {
			log.Fatalf("Error marshling the ticket: %d", http.StatusBadRequest)
			http.Error(w, "Error marshling the ticket:", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ticketJson)
		log.Print("Ticket found succesfully")
	}
}

func UpdateTicket(t postgres.TicketDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Update Handler by Id")
		if r.Method != http.MethodPut {
			log.Fatalf("Method not allowed %d", http.StatusBadRequest)
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
		id, err := getIdFromPath(r.URL.Path)
		if err != nil {
			log.Fatalf("Error getting id from Path: %d", http.StatusBadRequest)
			http.Error(w, "Error getting id from Path", http.StatusBadRequest)
		}
		var ticket ticket.Ticket
		if json.Decoder(r.Body).Decode(&ticket); err != nil {
			log.Fatalf("Error decoding body: %d", http.StatusBadRequest)
			http.Error(w, "Error decoding body", http.StatusBadRequest)
		}
		ticketUpdate, err := t.UpdateTicketDAO(r.Context(), ticket, id)
		if err != nil {
			// Here i can manage something bcs what if there's no ticket with the ID?
			log.Fatalf("Error updating the ticket: %d", http.StatusBadRequest)
			http.Error(w, "Error updating the ticket", http.StatusBadRequest)
		}
		ticketJson, err := json.Marshal(ticketUpdate)
		w.Header().Set("Content-Type", "application/json")
		w.Write(ticketJson)
		log.Print("Ticket update was succesfull")
	}
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
