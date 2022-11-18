package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.String("input", "", "user input from cmd")
	flag.Parse()

	if *input == "" {
		fmt.Println("Welcome to velotio")
		return
	}

	fmt.Println(*input)
}
