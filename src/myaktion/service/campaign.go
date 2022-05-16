package service

import (
	"github.com/RomanoLu/myaktion-go/src/myaktion/db"
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
	result := db.DB.Create(campaign)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Tracef("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	result := db.DB.Preload("Donations").Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaigns)
	return campaigns, nil
}
func GetCampaign(id uint) (*model.Campaign, error) {
	var campaign *model.Campaign
	result := db.DB.Preload("Donations").Find(&campaign, id)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaign)
	return campaign, nil
}

func UpdateCampaign(id uint, campaign *model.Campaign) (*model.Campaign, error) {
	existingCampaign, err := GetCampaign(id)
	if err != nil {
		return existingCampaign, err
	}
	db.DB.Preload("Donations").Model(&existingCampaign).Updates(campaign)
	entry := log.WithField("ID", id)
	entry.Info("Successfully updated campaign.")
	entry.Tracef("Updated: %v", campaign)
	return existingCampaign, nil
}

func DeleteCampaign(id uint) (*model.Campaign, error) {
	campaign, err := GetCampaign(id)
	if err != nil {
		return campaign, err
	}
	db.DB.Delete(campaign)
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted campaign.")
	entry.Tracef("Deleted: %v", campaign)
	return campaign, nil
}
