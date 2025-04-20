package main

import (
	"log"
	"net/http"

	"github.com/jakefurlong/flowwatch/backend/api"
)

func main() {
	http.HandleFunc("/api/river/current", api.GetCurrentRiverData)
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
