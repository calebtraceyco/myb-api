package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
)

//var (
//	initErrs []error
//	Port     string
//)

func initializeDatabase(cfg *config.Config, svc *facade.Service) error {
	psqlService, err := cfg.Database(PostgresDB)
	if err != nil {
		return err
	}
	svc.Db = psqlService.DB
	return nil
}

const PostgresDB = "PSQL"
