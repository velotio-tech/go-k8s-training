package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("-----------------------------------------------Welcome to Linux Shell-------------------------------------------")

	var inputCommand string

	for {
		fmt.Print("\n", getPrompt())
		fmt.Scanln(&inputCommand)

		inputCommand = strings.ToLower(inputCommand)

		if inputCommand == "ls" {
			listDirectoriesAndFiles()
		} else if inputCommand == "pwd" {
			fmt.Print(presentWorkingDirectory())
		} else if inputCommand == "exit" {
			os.Exit(0)
		}
	}
}
