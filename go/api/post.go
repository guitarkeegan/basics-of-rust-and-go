package api

import (
	"encoding/json"
	"net/http"

	"femm.com/server/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	var exhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&exhibition)
	if err != nil {
		// err.Error() will print the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Add(exhibition)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}
