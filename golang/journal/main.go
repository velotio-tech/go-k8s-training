package main

import (
	"github.com/velotio-ajaykumbhar/journal/app/service"
	"github.com/velotio-ajaykumbhar/journal/cmd"
	"github.com/velotio-ajaykumbhar/journal/util"
)

func main() {

	util.InitFile(service.AuthFilename)
	util.InitFile(service.SessionFilename)

	cmd.Execute()
}
