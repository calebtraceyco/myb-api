package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/dao/psql"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
)

func initializeDatabase(cfg *config.Config, svc *facade.Service) error {
	if psqlService, err := cfg.Database(PostgresDB); err != nil {
		return err
	} else {
		svc.PSQL = psql.DAO{Db: psqlService.DB}
	}
	return nil
}

const PostgresDB = "PSQL"
