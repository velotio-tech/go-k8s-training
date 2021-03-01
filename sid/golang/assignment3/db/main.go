package main

import (
	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
	dbserver "github.com/farkaskid/go-k8s-training/assignment3/db/server"
)

func main() {
	helpers.Bootstrap()
	dbserver.New(2222).ListenAndServe()
}
