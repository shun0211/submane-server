package main

import (
	"api/driver"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") == "development" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials/submane-firebase-adminsdk.json")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	driver.Init()
}
