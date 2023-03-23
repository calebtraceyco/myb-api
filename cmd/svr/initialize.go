package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
	"github.com/calebtraceyco/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infoln("initializing...")
}

type source struct{}

func (src source) Database(cfg *config.Config, svc *facade.Service) error {
	if psqlService, err := cfg.Database(PostgresDB); err != nil {
		return err
	} else {
		svc.PSQL = psql.DAO{Pool: psqlService.Pool}

	}
	return nil
}

const PostgresDB = "PSQL"
