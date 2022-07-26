package main

import (
	"ascan/desafio-go/database"
	"ascan/desafio-go/rest"
	"ascan/desafio-go/service"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting users service")

	db, err := database.NewDatabase()

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	service := service.NewService(db)
	rest := rest.NewRest(service)

	err = rest.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start service: %v", err)
	}
}