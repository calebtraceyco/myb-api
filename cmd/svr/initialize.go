package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
)

var service facade.Service
var initErrs []error

func init() {
	log.Infoln("=== initializing...")
	appConfig := config.New(configPath)

	initializeDatabase(appConfig)
}

func initializeDatabase(appConfig *config.Config) {
	if psqlService, err := appConfig.Database(PostgresDB); err != nil {
		initErrs = append(initErrs, err)
	} else {
		service.Db = psqlService.DB
	}
}

const PostgresDB = "PSQL"
