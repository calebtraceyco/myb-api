package user

import (
	"context"
	"github.com/calebtraceyco/mind-your-business-api/external"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
)

type DAOI interface {
	AddUser(ctx context.Context, user *models.User) (resp *external.ExecResponse, err error)
}

type DAO struct {
	PSQL   psql.DAOI
	Mapper MapperI
}

func (s DAO) AddUser(ctx context.Context, user *models.User) (resp *external.ExecResponse, err error) {

	if resp, err = s.PSQL.ExecContext(ctx, s.Mapper.MapUserExec(user)); err != nil {
		return nil, err
	}

	return resp, err
}
