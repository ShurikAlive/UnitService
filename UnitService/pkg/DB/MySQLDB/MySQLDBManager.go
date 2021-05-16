package MySQLDB

import (
	configs "UnitService/pkg/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sirupsen/logrus"
	"sync"
)

var instance *sql.DB
var once sync.Once

func ConnectDB() (*sql.DB) {

	params, err := configs.ParseEnv()
	db, err := sql.Open(params.DBType, params.DBUsername + ":" + params.DBPassword + "@/" + params.DBName + "?multiStatements=true")//MySQL80
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil
	}

	return db
}

func DisconectDB() {
	if instance != nil {
		instance.Close()
	}
}

func GetDBInstance() *sql.DB {
	once.Do(func() {
		instance = ConnectDB()
	})
	return instance
}

