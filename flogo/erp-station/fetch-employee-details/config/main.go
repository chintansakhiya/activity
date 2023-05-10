package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var AllConfig DBConfig

// GetConfig Collects all configs
func GetConfig() DBConfig {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err, "1")
	}
	err = godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err, "2")
	}
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println(err, "3")
	}

	AllConfig = DBConfig{}

	err = envconfig.Process("APP_PORT", &AllConfig)
	if err != nil {
		log.Fatal(err)
	}

	return AllConfig
}
