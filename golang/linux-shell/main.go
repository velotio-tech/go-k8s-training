package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-----------------------------------------------Welcome to Linux Shell-------------------------------------------\n")

	var inputCommand string

	for {
		fmt.Print("\n", getPrompt())
		fmt.Scanln(&inputCommand)

		inputCommand = strings.ToLower(inputCommand)

		if inputCommand == "ls" {
			listDirectoriesAndFiles()
		}
	}
}
