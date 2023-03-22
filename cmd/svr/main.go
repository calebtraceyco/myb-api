package main

import (
	"github.com/NYTimes/gziphandler"
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	"github.com/calebtracey/mind-your-business-api/internal/routes"
	log "github.com/sirupsen/logrus"
)

const configPath = "dev_config.yaml"

func main() {
	defer panicQuit()
	log.Infoln("initializing...")
	appConfig := config.New(configPath)
	appService := facade.Service{}

	if err := initializeDatabase(appConfig, &appService); err != nil {
		log.Errorf("failed to initialize database: %s", err)
		panicQuit()
	}

	log.Fatal(listenAndServe(appConfig.Port.Value, gziphandler.GzipHandler(
		corsHandler().Handler(
			routes.Handler{Service: appService}.Routes(),
		)),
	))
}

func panicQuit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
