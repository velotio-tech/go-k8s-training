package main

import "fmt"

type history []string

func (h history) printHistory() {
	for _, cmd := range h {
		fmt.Println(cmd)
	}
}
