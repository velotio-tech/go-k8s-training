package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Returns whether exit shell was called or not.
func processCommand(cmd string) bool {
	const exitShell = true
	cmdTokens := strings.Fields(cmd)
	if len(cmdTokens) == 0 {
		// Empty new line, no command.
		return !exitShell
	}
	switch cmdTokens[0] {
	case "ls":
		processListingCommand(cmdTokens[1:])
	case "pwd":
		showPresentWorkingDir()
	case "cd":
		changeWorkingDir(cmdTokens[1:])
	case "exit":
		return exitShell
	default:
		fmt.Println("command not found: ", cmdTokens[0])
	}
	return !exitShell
}

// linux ls command
func processListingCommand(args []string) {
	if len(args) == 0 {
		// Listing current directory
		currentDir, _ := os.Getwd()
		files, err := os.ReadDir(currentDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
	// else listing directory mentioned in args.
	// example: ls ~/some/path /another/path
	for _, path := range args {
		_, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		} else {
			files, _ := os.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
	}
}

func showPresentWorkingDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)
}

func changeWorkingDir(args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Invalid arguments")
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println(err)
	}
}
