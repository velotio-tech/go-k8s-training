package main

import (
	"flag"
	"fmt"
)

const (
	HELLO = "HELLO WORLD"
)

func main() {
	input := flag.String("input", HELLO, "Input Text")
	flag.Parse()
	fmt.Println(*input)
}
