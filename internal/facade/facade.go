package facade

import "database/sql"

type ServiceI interface {
}

type Service struct {
	Db *sql.DB
}
