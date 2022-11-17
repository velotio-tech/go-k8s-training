package main

import (
	"github.com/velotio-ajaykumbhar/microservice/user/api"
)

func main() {
	router := api.Setup()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
