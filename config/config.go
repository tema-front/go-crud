package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tema-front/go-crud/utils"
)

type Config struct {
	PORT  string
	DB_URL string
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
			log.Fatal("Couldn't load .env file")
	}
}

func LoadConfig() Config {
	port := utils.GetEnvValue("PORT")
	dbURL := utils.GetEnvValue("DB_URL")

	return Config{PORT: port, DB_URL: dbURL}
}
