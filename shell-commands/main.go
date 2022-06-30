package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

// To save history
var historyMap map[int]string

func main() {
	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get current username
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	username := currentUser.Username
	prefix := username + "@" + hostname + " "
	// Continue scanning of commands till user enters exit command
	for {
		// Get current working directory
		mydir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		workingDir := prefix + mydir + " % "
		fmt.Print(workingDir)
		// Scan command
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			break
		}
		exit := runCommand(text)
		if exit {
			break
		}
	}
}

func runCommand(command string) bool {
	fields := strings.Fields(command)
	switch fields[0] {
	case "ls":
		getListing()
	case "pwd":
		getWorkingDirectory()
	case "history":
		getHistory()
	case "cd":
		changeDirectory(fields)
	case "exit":
		return true
	default:
		fmt.Printf("%s : Invalid command\n", fields[0])
	}
	// Add command to history map
	addHistory(command)
	return false
}
