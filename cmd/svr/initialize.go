package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
)

var (
	appConfig  *config.Config
	appService facade.Service
	initErrs   []error
	Port       string
)

func init() {
	log.Infoln("=== initializing...")
	appConfig = config.New(configPath)
	appService = facade.Service{}
	Port = appConfig.Port.Value

	initializeDatabase()
}

func initializeDatabase() {
	if psqlService, err := appConfig.Database(PostgresDB); err != nil {
		initErrs = append(initErrs, err)
	} else {
		appService.Db = psqlService.DB
	}
}

const PostgresDB = "PSQL"
