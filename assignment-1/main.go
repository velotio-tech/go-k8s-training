package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

// pastCommands stores all the commands executed during the current session
var pastCommands []string

// constants declared for the commands
const (
	cd      = "cd"
	ls      = "ls"
	pwd     = "pwd"
	history = "history"
)

func main() {
	fmt.Print(time.Now().Format(time.RFC1123))
	fmt.Println(" Session Starts...")

	for {
		printCurrentDir()

		command, arg := takeUserInput()

		if command == "exit" {
			break
		}
		processCommands(command, arg)
	}

	fmt.Print(time.Now().Format(time.RFC1123))
	fmt.Println(" Session Ends...")
}

func processCommands(command, arg string) {
	currentPath := currentDirPath()
	switch command {
	case "":
		break
	case ls:
		printDirInfo(currentPath, arg)
	case cd:
		changeDir(currentPath, arg)
	case pwd:
		fmt.Println(currentPath)
	case history:
		printHistory()
	default:
		fmt.Println(command + ": command not found")
	}
}

// changeDir: Implementation for cd command
func changeDir(currPath, dirPath string) {
	err := os.Chdir(currPath + dirPath)
	checkErr(err)
}

// currentDirPath: Implementation for pwd command
func currentDirPath() string {
	path, err := os.Getwd()
	checkErr(err)
	path += "/"
	return path
}

// printDirInfo: Implementation for ls command
func printDirInfo(currPath, dirPath string) {
	files, err := os.ReadDir(currPath + dirPath)
	checkErr(err)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// takeUserInput: For processing the user input
func takeUserInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	checkErr(err)
	userInput := scanner.Text()
	userInput = strings.Trim(userInput, " ")
	inputs := strings.Split(userInput, " ")
	command := ""
	arg := ""
	if len(inputs) > 0 {
		command = inputs[0]
	}
	if len(inputs) > 1 {
		arg = inputs[1]
	}
	pastCommands = append(pastCommands, userInput)
	return command, arg
}

// printHistory: for printing the history or past commands
func printHistory() {
	for ind, command := range pastCommands {
		fmt.Println(ind+1, command)
	}
}

// printCurrentDir: for printing the username, hostname and curr path
func printCurrentDir() {
	path, err := os.Getwd()
	checkErr(err)
	hostName, err := os.Hostname()
	checkErr(err)
	user, err := user.Current()
	checkErr(err)
	userName := user.Username
	rootPath := userName + "@" + hostName + " " + path
	fmt.Printf("%s$ ", rootPath)
}

// logging errors
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
