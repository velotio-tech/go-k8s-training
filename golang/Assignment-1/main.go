package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	history := history{}

	fmt.Println("__________Welcome to Interactive Linux Shell___________")
	prompt := getPrompt()
	fmt.Println(prompt)
	for {
		reader := bufio.NewReader(os.Stdin)
		inputCommand, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		inputCommand = strings.TrimSuffix(inputCommand, "\n")
		history = append(history, inputCommand)
		runCommand(inputCommand, history)
	}
}

func runCommand(inputCommand string, h history) {
	tokens := strings.Fields(inputCommand)
	command, arguments := tokens[0], tokens[1:]
	switch command {
	case "exit":
		exitShell()
	case "pwd":
		pwd := getCurrentDirectory()
		fmt.Println(pwd)
	case "ls":
		response := getListOfFilesAndDirectories()
		for _, name := range response {
			fmt.Println(name)
		}
	case "cd":
		changeDirectory(arguments[0])
	case "history":
		h.printHistory()
	default:
		fmt.Println("Invalid command, try again")
	}
}
