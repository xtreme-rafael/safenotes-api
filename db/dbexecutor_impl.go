package db

import (
	"database/sql"
	"log"
	"errors"
)

type DBExecutorImpl struct {
	db *sql.DB
}

func NewDBExecutor(database *sql.DB) DBExecutor {
	if database == nil {
		log.Panic("Can't create DB executor with nil DB")
	}

	return &DBExecutorImpl{db: database}
}

func (self *DBExecutorImpl) ExecWithDB(worker func(*sql.DB) error) error {
	return errors.New("not implemented")
}

func (self *DBExecutorImpl) ExecWithTx(worker func(*sql.Tx) error) error {
	return errors.New("not implemented")
}

func (self *DBExecutorImpl) CloseDB() error {
	return self.db.Close()
}
