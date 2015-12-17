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
	PackageName = "DB"
	mysqlDataSourceStringFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=true"
)

func GetCloudFoundryDatabase() (*sql.DB, error) {
	dataSourceName := getCloudFoundryDataSource()
	if dataSourceName == "" {
		err := errors.New("Unable to retrieve database credentials from CF")
		utils.Log(PackageName, err.Error())
		return nil, err
	}

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		err = errors.New("Unable to open connection with DB - " + err.Error())
		utils.Log(PackageName, err.Error())
		return nil, err
	}

	return db, nil
}

func getCloudFoundryDataSource() string {
	env, _ := cfenv.Current()
	if env != nil {
		mysqlService, _ := env.Services.WithTag("mysql")
		credentials := mysqlService[0].Credentials
		dataSourceName := fmt.Sprintf(mysqlDataSourceStringFormat,
			credentials["username"],
			credentials["password"],
			credentials["hostname"],
			credentials["port"],
			credentials["name"])
		return dataSourceName
	}

	return ""
}
