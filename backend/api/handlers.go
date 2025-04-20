package api

import (
	"encoding/json"
	"net/http"

	"github.com/jakefurlong/flowwatch/backend/fetcher"
)

func GetCurrentRiverData(w http.ResponseWriter, r *http.Request) {
	site := r.URL.Query().Get("site")
	if site == "" {
		http.Error(w, "Missing 'site' query param", http.StatusBadRequest)
		return
	}

	data, err := fetcher.FetchRiverData(site)
	if err != nil {
		http.Error(w, "Error fetching data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
