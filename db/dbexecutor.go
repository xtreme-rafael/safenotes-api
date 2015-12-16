package db

import (
	"database/sql"
)

type DBExecutor interface {
	ExecWithDB(func(*sql.DB) error) error
	ExecWithTx(func(*sql.Tx) error) error
}
