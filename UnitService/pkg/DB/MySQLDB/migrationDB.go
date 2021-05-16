package MySQLDB

import(
	"UnitService/pkg/config"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func MakeMigrationDB() (error) {
	driver, err := mysql.WithInstance(GetDBInstance(), &mysql.Config{})
	if err != nil {
		return err
	}
	conf, err := config.ParseEnv()
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		conf.DBMigrationsPath,
		conf.DBName,
		driver,
	)
	if err != nil {
		return err
	}

	m.Up();//!!! err

	return nil
}