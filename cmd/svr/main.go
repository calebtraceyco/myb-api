package main

import (
	"github.com/NYTimes/gziphandler"
	"github.com/calebtracey/mind-your-business-api/internal/routes"
	log "github.com/sirupsen/logrus"
)

const configPath = "dev_config.yaml"

func main() {

	if initErrs != nil {
		log.Error(initErrs)
		panicQuit()

	} else {

		log.Fatal(listenAndServe("8080", gziphandler.GzipHandler(
			corsHandler().Handler(
				routes.Handler{Service: service}.
					Routes(),
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
