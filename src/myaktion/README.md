# curls

# createCampaign
curl -H "Content-Type: application/json" -d '{"name":"Trikots","donationMinimum":2,"targetAmount":100,"account":{"name":"Martin","bankName":"DKB","number":"123456"}}' localhost:8000/campaign

# updateCampaign
curl -H "Content-Type: application/json" -d '{"name":"Trikots für BFC","donationMinimum":20,"targetAmount":200,"account":{"name":"Luca","bankName":"DKB","number":"123456"}}' localhost:8000/campaign/1

# AddDonation
curl -H "Content-Type: application/json" -d'{"campaignID":1,"amount":200,"donorName":"Luca","receiptRequested":true, "status": "TRANSFERRED", "account":{"name":"Luca","bankName":"DKB","number":"793"}}' localhost:8000/campaigns/1/donation