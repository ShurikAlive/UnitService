package DB

import (
	"database/sql"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sirupsen/logrus"
)

var instance *sql.DB
var once sync.Once

func ConnectDB() (*sql.DB) {

	//params, err := parseEnv()
	db, err := sql.Open("mysql", "root:Future1994!)@/cafe_test")//MySQL80
	if err != nil {
		return nil
	}
	//defer db.Close()

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil
	}

	return db
}

func GetDBInstance() *sql.DB {
	once.Do(func() {
		instance = ConnectDB()
	})
	return instance
}

