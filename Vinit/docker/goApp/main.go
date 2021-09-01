package main

import (
	"flag"
	"fmt"
)
func main() {
	inp := flag.String("input", "Hello World!", "Trying the cli flags")
	flag.Parse()
	fmt.Println(*inp)
}
