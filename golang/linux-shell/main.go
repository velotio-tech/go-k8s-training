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

	for {
		fmt.Print(getPrompt())
		scanner.Scan()

		tokens := strings.Fields(scanner.Text())
		command, arguments := tokens[0], tokens[1:]

		switch command {
		case "ls":
			printFilesAndDirectories()
		case "pwd":
			fmt.Println(presentWorkingDirectory())
		case "exit":
			fmt.Println("Shutting down the shell......")
			os.Exit(0)
		case "cd":
			os.Chdir(arguments[0])
		}
	}
}
