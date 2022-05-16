package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RomanoLu/myaktion-go/src/myaktion/model"
	"github.com/RomanoLu/myaktion-go/src/myaktion/service"
	log "github.com/sirupsen/logrus"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Error("Can't serialize request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(&campaign); err != nil {
		log.Error("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaign)
}

func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Errorf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, campaigns)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	log.Info(id)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaigns, err := service.GetCampaign(id)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, campaigns)
}

func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Error("Can't serialize request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaigns, err := service.UpdateCampaign(campaign, id)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, campaigns)
}

func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.DeleteCampaign(id)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func AddDonation(w http.ResponseWriter, r *http.Request) {
	var donation model.Donaition
	err := json.NewDecoder(r.Body).Decode(&donation)
	if err != nil {
		log.Error("Can't serialize request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := service.AddDonation(donation,id); err != nil {
		log.Errorf("Error calling service GetCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

