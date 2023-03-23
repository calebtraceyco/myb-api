package user

import (
	"context"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
)

type DAOI interface {
	NewUser(ctx context.Context, params any) (resp any, errs []error)
}

type DAO struct {
	PSQL psql.DAOI
}

// NewUser TODO probably dont need this, can just call what is in here inside of the facade function
func (s DAO) NewUser(ctx context.Context, params any) (resp any, errs []error) {
	var err error
	exec := ""

	// TODO map params to psql query

	if resp, err = s.PSQL.ExecContext(ctx, exec); err != nil {
		return resp, []error{err}
	}

	// TODO mapping stuff here

	return resp, nil
}
