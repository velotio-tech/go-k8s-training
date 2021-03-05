// main.go

package main

import "os"

func main() {
	// create new app
	a := App{}
	// initialize the app and run it
	a.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_PORT"))

	a.Run(":8010")
}
