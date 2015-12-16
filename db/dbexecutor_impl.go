package db

import (
	"database/sql"
	"log"
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
	return worker(self.db)
}

func (self *DBExecutorImpl) ExecWithTx(worker func(*sql.Tx) error) error {
	tx, err := self.db.Begin()
	if err != nil {
		log.Println("failed to create DB transaction: ", err.Error())
		return err
	}

	err = worker(tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func (self *DBExecutorImpl) CloseDB() error {
	return self.db.Close()
}
