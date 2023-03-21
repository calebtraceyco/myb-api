package main

import (
	"github.com/NYTimes/gziphandler"
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business/internal/routes"
	log "github.com/sirupsen/logrus"
)

const configPath = "local_config.yaml"

func main() {
	log.Infoln("=== Initializing...")

	if svc, errs := initializeDAO(config.New(configPath)); errs != nil {
		log.Error(errs)
		panicQuit()

	} else {

		log.Fatal(listenAndServe("8080", gziphandler.GzipHandler(
			corsHandler().Handler(
				routes.Handler{
					Service: svc,
				}.Routes(),
			)),
		))
	}
}

func panicQuit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
