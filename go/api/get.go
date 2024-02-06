package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"femm.com/server/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// should check that it is a POST request here
	w.Header().Set("Content-Type", "application/json")
	// api/exhibitions/?id=420
	// or id := r.URL.Query()["id"] would return a slice
	id := r.URL.Query().Get("id")
	if id != "" {
		finalId, err := strconv.Atoi(id)
		// GetAll() should be changed to getOne()
		if err == nil {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
