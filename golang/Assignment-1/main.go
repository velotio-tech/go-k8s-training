package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("__________Welcome to Interactive Linux Shell___________")
	prompt := getPrompt()
	fmt.Println(prompt)
	for {
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		command = strings.TrimSuffix(command, "\n")
		runCommand(command)
	}
}

func runCommand(command string) {
	if command == "exit" {
		exitShell()
	} else if command == "pwd" {
		pwd := getCurrentDirectory()
		fmt.Println(pwd)
	} else if command == "ls" {
		response := getListOfFilesAndDirectories()
		fmt.Println(response)
	} else {
		fmt.Println("Invalid command, try again")
	}
}
