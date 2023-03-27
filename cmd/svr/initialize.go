package main

import (
	cfg "github.com/calebtraceyco/config"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/user"
	"github.com/calebtraceyco/mind-your-business-api/internal/facade"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Infoln("initializing...")
}

type source struct{}

func (src source) Database(cfg *cfg.Config, svc *facade.Service) error {
	if psqlService, err := cfg.Database(PostgresDB); err != nil {
		return err
	} else {
		svc.UserDAO = user.DAO{
			PSQL:   psql.DAO{Pool: psqlService.Pool},
			Mapper: user.Mapper{},
		}
	}
	return nil
}

const PostgresDB = "PSQL"
