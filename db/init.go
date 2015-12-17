package db

const (
	sqlDialectMysql   = "mysql"
	cfServiceTagMysql = "mysql"
)

var executor DBExecutor

func InitDB() error {
	return initExecutor()
}

func GetExecutor() DBExecutor {
	return executor
}

func initExecutor() error {
	db, err := GetCloudFoundryDatabase(cfServiceTagMysql, sqlDialectMysql)

	if err != nil {
		return err
	}

	if err = MigrateDB(db, sqlDialectMysql); err != nil {
		return err
	}

	executor = NewDBExecutor(db)
	return nil
}
