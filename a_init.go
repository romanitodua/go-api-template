package main

import (
	"log"
	"os"
)

func verifyEnvs() {
	requiredEnvs := []string{
		"Origins", "DB_CONNECTION",
	}
	for _, env := range requiredEnvs {
		present := os.Getenv(env)
		if present != "" {
			log.Fatalf("Missing required environment variable: %s", env)
		}
	}
	log.Println("All environment variables are present")
}
