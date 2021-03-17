package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/farkaskid/go_assignments/assignment1/helpers"
)

func Handle(cmd string, history *helpers.History) int {
	// Main handler for commands, this will branch off to different handlers for different commands

	commandFields := strings.Fields(cmd)

	switch commandFields[0] {

	case "ls":
		return lsHandler(commandFields[1:])

	case "cd":
		return cdHander(commandFields[1:])

	case "pwd":
		return pwdHandler(commandFields[1:])

	case "history":
		return historyHandler(history)

	default:
		fmt.Println("Unknowm command:", cmd)
		return 1
	}
}

func lsHandler(inputs []string) int {
	// Handles the ls command

	currentDir, err := os.Open(".")

	if err != nil {
		helpers.PrintError("Failed to open current directory", err)

		return 1
	}

	defer currentDir.Close()

	files, err := currentDir.Readdirnames(0)

	if err != nil {
		helpers.PrintError("Failed to read current directory", err)

		return 1
	}

	for _, file := range files {
		fmt.Println(file)
	}

	return 0
}

func cdHander(input []string) int {
	// Handler for the 'cd' command

	dst := input[0]

	if strings.HasPrefix(dst, "~") {
		dst = strings.Replace(dst, "~", os.Getenv("HOME"), 1)
	}

	err := os.Chdir(dst)

	if err != nil {
		helpers.PrintError("Failed to change the directory", err)

		return 1
	}

	return 0
}

func pwdHandler(input []string) int {
	// Handler for the "pwd" command

	cwd := helpers.Getcwd()

	if len(cwd) == 0 {
		return 1
	}

	fmt.Println(cwd)

	return 0
}

func historyHandler(history *helpers.History) int {
	// Prints the history

	for elem := history.Commands.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value)
	}

	return 0
}
