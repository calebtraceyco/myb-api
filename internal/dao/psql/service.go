package psql

import (
	"context"
	"database/sql"
)

type DAOI interface {
	ExecContext(ctx context.Context, exec string) (resp any, err error)
}

type DAO struct {
	Db *sql.DB
}

func (s DAO) ExecContext(ctx context.Context, exec string) (resp any, err error) {
	if resp, err = s.Db.ExecContext(ctx, exec); err != nil {
		return nil, err
	}

	return resp, nil
}
