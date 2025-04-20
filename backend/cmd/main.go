package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jakefurlong/flowwatch/backend/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/river/current", api.GetCurrentRiverData).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
