package utils

import (
	"log"
	"os"
)

func GetEnvValue(name string) string {
	value := os.Getenv(name)
	
	if value == "" {
		log.Fatalf("%v is not found in environment", name)
	} else {
		log.Printf("%v has been successfully found", name)
	}

	return value
}