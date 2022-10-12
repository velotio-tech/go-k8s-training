package main

import (
	"fmt"
	"log"

	"github.com/velotio-tech/go-k8s-training/shell"
)

func main() {
	if err := run(); err != nil {
		log.Panic(err)
	}
}

func run() error {
	shell, err := shell.New()
	if err != nil {
		return err
	}
	for {
		err := shell.Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}
