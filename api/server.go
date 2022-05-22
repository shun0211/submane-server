package main

import (
	"api/driver"
	"os"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials/submane-firebase-adminsdk.json")
	driver.Init()
}
