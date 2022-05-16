package main

import (
	"net/http"

	"github.com/RomanoLu/myaktion-go/src/myaktion/db"
	"github.com/RomanoLu/myaktion-go/src/myaktion/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/campaign", handler.CreateCampaign).Methods("POST")
	router.HandleFunc("/health", handler.Health).Methods("GET")
	router.HandleFunc("/campaigns", handler.GetCampaigns).Methods("GET")
	router.HandleFunc("/campaigns/{id}", handler.GetCampaign).Methods("GET")
	router.HandleFunc("/campaigns/{id}", handler.UpdateCampaign).Methods("UPDATE")
	router.HandleFunc("/campaigns/{id}", handler.DeleteCampaign).Methods("DELETE")
	router.HandleFunc("/campaigns/{id}/donation", handler.AddDonation).Methods("POST")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}

func init() {
	db.Init()
	// init logger
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	/*level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)*/
}
