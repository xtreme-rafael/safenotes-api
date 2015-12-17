package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rubenv/sql-migrate"
	"github.com/xtreme-rafael/safenotes-api/db/migrations"
	"github.com/xtreme-rafael/safenotes-api/utils"
)

const (
	migrationsDir = "db/migrations" // NOTE: generated from db/sql folder with go-bindata
)

func MigrateDB(db *sql.DB, sqlDialect string) error {
	if db == nil {
		err := errors.New("Error running migrations - DB is nil")
		utils.Log(PackageName, err.Error())
		return err
	}

	utils.Log(PackageName, "Starting DB migrations")

	migrationSrc := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      migrationsDir,
	}

	n, err := migrate.Exec(db, sqlDialect, migrationSrc, migrate.Up)
	if err != nil {
		err = fmt.Errorf("Error running migrations - %s", err.Error())
		utils.Log(PackageName, err.Error())
		return err
	}

	utils.Log(PackageName, fmt.Sprintf("Completed %d DB migrations", n))
	return nil
}
