package main

import (
	"log"
	"net/http"

	"github.com/JoaoGumiero/Crud_Backend/api"
)

func main() {
	mux := http.NewServeMux()
	api.UploadRoutes(mux)

	log.Println("Listening on port :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
