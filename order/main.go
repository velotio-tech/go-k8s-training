package main

import (
	"github.com/velotio-tech/go-k8s-training/order/controller"
	"github.com/velotio-tech/go-k8s-training/order/utils"
)

func main() {
	utils.InitConfig()
	controller.NewController().Run()
}
