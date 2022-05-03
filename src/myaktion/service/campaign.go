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