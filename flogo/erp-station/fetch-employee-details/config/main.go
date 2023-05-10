package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var AllConfig DBConfig

// GetConfig Collects all configs
func GetConfig() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	AllConfig = DBConfig{}

	err = envconfig.Process("APP_PORT", &AllConfig)
	if err != nil {
		log.Fatal(err)
	}

	return AllConfig
}
