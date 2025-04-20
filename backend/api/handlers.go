package api

import (
	"encoding/json"
	"net/http"

	"flowwatch/fetcher"
)

func GetCurrentRiverData(w http.ResponseWriter, r *http.Request) {
	site := r.URL.Query().Get("site")
	if site == "" {
		http.Error(w, "Missing site parameter", http.StatusBadRequest)
		return
	}

	data, err := fetcher.GetRiverData(site)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
