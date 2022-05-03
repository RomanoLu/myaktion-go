package handler

import (
	"encoding/json"
	"github.com/RomanoLu/myaktion-go/src/myaktion/model"
	"github.com/RomanoLu/myaktion-go/src/myaktion/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Info("Can't serialize request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(&campaign); err != nil {
		log.Info("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaign); err != nil {
		log.Info("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Info("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaigns); err != nil {
		log.Info("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
