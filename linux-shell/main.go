package main

import (
	"bufio"
	"fmt"
	"os"
)

// created a map to save history of commands executed
var history = make(map[int]string)

// created a map to save collection of valid commands with info
var validCMD = map[string]string{
	"ls": "used to list the contents of the directory",
	"pwd": "used to prints the path of the working directory",
	"history": "used to view the previously executed commands",
	"cd": "used to change current working directory",
	"help": "used to view all available commands",
	"exit": "used to exit the shell",
}

func main() {
	
	// yellow color for pretext and blue for other text
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"

	for {

		// prefix consists of username + hostname + current working directory
		prefix := getPrefix()
		
		fmt.Print(string(colorYellow),prefix)
		fmt.Print(string(colorBlue))

		// to read input from standard input
		reader := bufio.NewReader(os.Stdin)

		// read until new line
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			break
		}
		
		// add the command to history
		addToHistory(input, history)
		
		// execute the entered command
		executeCMD(input, history, validCMD)
	}
}