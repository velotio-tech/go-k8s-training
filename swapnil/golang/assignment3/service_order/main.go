// main.go

package main

import "os"

func main() {
	// create a App instance, get database config from env and run the service
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_PORT"))

	a.Run(":8011")
}
