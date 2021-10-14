package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func showPrompt() {
	currentUser, _ := user.Current()
	pwd, _ := os.Getwd()
	hostName, _ := os.Hostname()
	fmt.Printf("%v@%v %v ", currentUser.Username, hostName, pwd)
}

func runShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		showPrompt()
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		command := scanner.Text()

		isExit := processCommand(command)
		if isExit {
			break
		}
	}
}
