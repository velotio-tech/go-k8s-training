package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {

	fmt.Println("-----------------Welcome to Linux Shell-------------")

	scanner := bufio.NewScanner(os.Stdin)

	history := restoreHistory()

	for {

		fmt.Println(getPromt())
		scanner.Scan()

		inputCommand := scanner.Text()
		tokens := strings.Fields(inputCommand)
		command, arguments := tokens[0], tokens[1:]

		switch command {
		case "ls":
			listDirectoriesAndFiles()
		case "pwd":
			presentWorkingDir()
		case "exit":
			fmt.Println("Shutting down the shell......")
			history.update(inputCommand)
			os.Exit(0)
		case "cd":
			os.Chdir(arguments[0])
		case "history":
			history.display()
		default:
			fmt.Println("Command not found:", command)
		}
		history.update(inputCommand)
	}
}

func presentWorkingDir() {

	currentDirPath, _ := os.Getwd()

	fmt.Print(currentDirPath)

}

func listDirectoriesAndFiles() {
	files, _ := os.ReadDir(".")

	for _, file := range files {
		fmt.Print(file.Name(), "\t")
	}
}

func getPromt() string {
	currUser, _ := user.Current()
	hostname, _ := os.Hostname()
	return "[" + currUser.Username + "@" + hostname + "]:" + getCurrentDirectory() + " $ "
}

func getCurrentDirectory() string {
	currentDirPath, _ := os.Getwd()

	directories := strings.Split(currentDirPath, "/")

	return directories[len(directories)-1]
}
