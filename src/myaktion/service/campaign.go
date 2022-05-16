package service

import (
	"github.com/RomanoLu/myaktion-go/src/myaktion/model"
	log "github.com/sirupsen/logrus"
)

var (
	campaignStore map[uint]*model.Campaign
	actCampaignId uint = 1
)

func init() {
	campaignStore = make(map[uint]*model.Campaign)
}
func CreateCampaign(campaign *model.Campaign) error {
	campaign.ID = actCampaignId
	campaignStore[actCampaignId] = campaign
	actCampaignId += 1
	log.Info("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Trace("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	for _, campaign := range campaignStore {
	campaigns = append(campaigns, *campaign)
	}
	log.Trace("Retrieved: %v", campaigns)
	return campaigns, nil
}

func GetCampaign(id uint) (model.Campaign, error) {
	var requestedCampaign model.Campaign
	requestedCampaign = *campaignStore[id]
	log.Trace("Retrieved: %v", requestedCampaign)
	return requestedCampaign, nil
}

func UpdateCampaign(campaign model.Campaign, id uint) (model.Campaign, error) {
	campaignStore[id] = &campaign
	log.Trace("Updated: %v", campaign)
	return campaign, nil
}

func DeleteCampaign(id uint) error{
	delete(campaignStore, id)
	log.Trace("Deleted: %v", id)
	return nil
}

func AddDonation(donation model.Donaition, id uint) error{
	 campaign := campaignStore[id]	
	 campaign.Donaitions = append(campaign.Donaitions, donation)
	 log.Trace("Added Donation: %v", donation)
	return nil
}

	