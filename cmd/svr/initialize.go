package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/dao/psql"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infoln("initializing...")
}

type InitializerI interface {
	Router(cfg *config.Config, svc *facade.Service) error
	Database(cfg *config.Config, svc *facade.Service) error
}
type Initializer struct{}

func (i *Initializer) Router(cfg *config.Config, svc *facade.Service) error {
	log.Infoln("Router...")
	return nil
}

func (i *Initializer) Database(cfg *config.Config, svc *facade.Service) error {
	if psqlService, err := cfg.Database(PostgresDB); err != nil {
		return err
	} else {
		svc.PSQL = psql.DAO{Db: psqlService.DB}
	}
	return nil
}

const PostgresDB = "PSQL"
