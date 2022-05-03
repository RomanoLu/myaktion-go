package main

import (
	"log"
	"net/http"
	"github.com/RomanoLu/myaktion-go/src/myaktion/handler"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET")
	if err := http.ListenAndServe(":8000", router); err != nil {
	log.Fatal(err)
	}
	}