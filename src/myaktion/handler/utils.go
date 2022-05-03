package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type result struct {
	Success string `json:"success"`
}

func sendJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
