package main

import (
	"github.com/velotio-tech/go-k8s-training/user/controller"
	"github.com/velotio-tech/go-k8s-training/user/utils"
)

func main() {
	utils.InitConfig()
	controller.NewController().Run()
}
