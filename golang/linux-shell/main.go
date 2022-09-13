package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("-----------------------------------------------Welcome to Linux Shell-------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)

	history := restoreHistory()

	for {
		fmt.Print(getPrompt())
		scanner.Scan()

		fullCommand := scanner.Text()
		tokens := strings.Fields(fullCommand)
		command, arguments := tokens[0], tokens[1:]

		switch command {
		case "ls":
			printFilesAndDirectories()
		case "pwd":
			fmt.Println(presentWorkingDirectory())
		case "exit":
			fmt.Println("Shutting down the shell......")
			history.update(fullCommand)
			os.Exit(0)
		case "cd":
			os.Chdir(arguments[0])
		case "history":
			history.display()
		default:
			fmt.Println("Command not found:", command)
		}

		history.update(fullCommand)
	}
}
