package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xtreme-rafael/safenotes-api/utils"
)

const (
	PackageName            = "DB"
	dataSourceStringFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=true"
)

func GetCloudFoundryDatabase(cfServiceTag string, sqlDialect string) (*sql.DB, error) {
	dataSourceName := getCloudFoundryDataSource(cfServiceTag)
	if dataSourceName == "" {
		err := errors.New("Unable to retrieve database credentials from CF")
		utils.Log(PackageName, err.Error())
		return nil, err
	}

	db, err := sql.Open(sqlDialect, dataSourceName)
	if err != nil {
		err = errors.New("Unable to open connection with DB - " + err.Error())
		utils.Log(PackageName, err.Error())
		return nil, err
	}

	return db, nil
}

func getCloudFoundryDataSource(cfServiceTag string) string {
	env, _ := cfenv.Current()
	if env != nil {
		service, _ := env.Services.WithTag(cfServiceTag)
		credentials := service[0].Credentials
		dataSourceName := fmt.Sprintf(dataSourceStringFormat,
			credentials["username"],
			credentials["password"],
			credentials["hostname"],
			credentials["port"],
			credentials["name"])
		return dataSourceName
	}

	return ""
}
