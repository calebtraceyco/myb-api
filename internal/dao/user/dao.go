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
	if exec, execErr := s.Mapper.MapUserExec(user); execErr == nil {
		if resp, err = s.PSQL.ExecContext(ctx, exec); err != nil {
			return nil, err
		}
	} else {
		return nil, execErr
	}
	return resp, nil
}

func (s DAO) GetUser(ctx context.Context) {}
