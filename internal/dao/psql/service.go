package psql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DAOI interface {
	ExecContext(ctx context.Context, exec string) (resp any, err error)
}

type DAO struct {
	Pool *pgxpool.Pool
}

func (s DAO) ExecContext(ctx context.Context, exec string) (resp any, err error) {
	if resp, err = s.Pool.Exec(ctx, exec); err != nil {
		return nil, err
	}

	return resp, nil
}
