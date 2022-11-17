package main

import (
	"github.com/velotio-ajaykumbhar/microservice/order/api"
)

func main() {
	router := api.Setup()

	if err := router.Run(":9090"); err != nil {
		panic(err)
	}
}
