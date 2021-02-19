package config

import (
	"os"
)

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type config struct {
	PostgresConfig *postgresConfig
}

var configInstance *config = nil

func GetConfig() *config {
	if configInstance != nil {
		return configInstance
	}

	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgUser := os.Getenv("DB_USERNAME")
	pgPass := os.Getenv("DB_PASSWORD")
	pgDB := os.Getenv("DB_NAME")

	return &config{
		PostgresConfig: &postgresConfig{
			Host:     pgHost,
			Port:     pgPort,
			User:     pgUser,
			Password: pgPass,
			DBName:   pgDB,
		},
	}
}
