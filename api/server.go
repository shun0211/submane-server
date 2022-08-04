package main

import (
	"api/driver"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	driver.Init()
}
