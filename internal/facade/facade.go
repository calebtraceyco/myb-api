package facade

import (
	"context"
	"github.com/calebtracey/mind-your-business-api/internal/dao/psql"
)

type ServiceI interface {
	NewUser(ctx context.Context, params any) (resp any, err error)
}

type Service struct {
	PSQL psql.DAOI
}

func (s Service) NewUser(ctx context.Context, params any) (resp any, err error) {

	// TODO add request validation
	// TODO parse params and map request query

	if resp, err = s.PSQL.ExecContext(ctx, ""); err != nil {
		return nil, err
	}

	// TODO add response mapping

	return resp, nil
}
