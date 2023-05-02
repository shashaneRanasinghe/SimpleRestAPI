package config

import (
	"github.com/joho/godotenv"
	"github.com/tryfix/log"
)

func LoadConfigs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading configs", err)
	}
}

type DBConnection struct {
	DBHost    string
	DBPort    string
	Username  string
	Password  string
	DBName    string
	DBNetwork string
}
