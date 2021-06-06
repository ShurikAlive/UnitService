package main

import "github.com/kelseyhightower/envconfig"

const appID = "UnitService"

type Config struct {
	ServeRESTAddress string `envconfig:"servr_rest_address" default:":8181"`

	DBType           string `envconfig:"database_type" default:"mysql"`
	DBName           string `envconfig:"database_name" default:"unit_db"`
	DBUsername       string `envconfig:"database_username" default:"root"`
	DBPassword       string `envconfig:"database_password" default:"Future1994!)"`
	DBMigrationsPath string `envconfig:"database_migrations_path" default:"file://db/migrations"`
}

func ParseEnv() (*Config, error) {
	c := new(Config)
	if err := envconfig.Process(appID, c); err != nil {
		return nil, err
	}
	return c, nil
}
