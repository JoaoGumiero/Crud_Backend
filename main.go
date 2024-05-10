package main

import (
	"context"
	"log"
	"net/http"

	"github.com/JoaoGumiero/Crud_Backend/api"
	"github.com/jackc/pgx/v5/pgxpool" // Concurrency safe
)

func main() {
	// the connection string for the postgres data base
	// Afterwards when i publish this repo, I'll need to fake those db access datas.
	connString := "postgres://postgres:123@localhost:5432/postgres"

	Dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	// Ping the database with a fake sql query to check if it responds with no issues, if it does, it'll catch an error.
	err = Dbpool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping the database: %v", err)
	}
	// Defer: The following function will only be excecuted when the adjcent function has a return or's stalled.
	defer Dbpool.Close()

	// New server and uploading the routes.
	mux := http.NewServeMux()
	api.UploadRoutes(mux)

	log.Println("Listening on port :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
