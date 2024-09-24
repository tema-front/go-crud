package utils

import (
	"log"
	"os"
)

func GetEnvValue(name string) string {
	value := os.Getenv(name)
	
	if value == "" {
		log.Fatalf("%v is not found in .env file", name)
	}

	return value
}