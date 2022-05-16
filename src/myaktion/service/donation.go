package service

import (
	"github.com/RomanoLu/myaktion-go/src/myaktion/db"
	"github.com/RomanoLu/myaktion-go/src/myaktion/model"
	log "github.com/sirupsen/logrus"
)

func AddDonation(campaignId uint, donation *model.Donation) error {
	donation.CampaignID = campaignId
	result := db.DB.Create(donation)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new campaign with ID %v in database.", donation.ID)
	log.Tracef("Stored: %v", donation)
	return nil
}