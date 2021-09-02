package main

import (
	"flag"
	"fmt"
)

func main() {
	msg := flag.String("input", "Hello World", "cli input for docker app")
	flag.Parse()
	fmt.Println(*msg)
}
