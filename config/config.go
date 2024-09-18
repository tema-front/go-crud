package config

import (
	"github.com/tema-front/go-crud/utils"
)

type Config struct {
	PORT  string
	DB_URL string
}

func LoadConfig() Config {
	port := utils.GetEnvValue("PORT")
	dbURL := utils.GetEnvValue("DB_URL")

	return Config{PORT: port, DB_URL: dbURL}
}
