package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func defaultPath() {
	currentUser, _ := user.Current()
	pwd, _ := os.Getwd()
	hostName, _ := os.Hostname()
	fmt.Printf("%v@%v %v ", currentUser.Username, hostName, pwd)
}

func runShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		defaultPath()
		// take input
		scanner.Scan()
		// execute the command
		isExit := executeCmd(scanner.Text())
		if isExit {
			break
		}
	}
}

func executeCmd(command string) bool {
	const exitShell = true
	cmd := strings.Fields(command)
	if len(cmd) == 0 {
		// no command
		return !exitShell
	}
	switch cmd[0] {
	case "ls":
		handleLs(cmd[1:])
	case "pwd":
		handlePwd()
	case "cd":
		handleCd(cmd[1:])
	case "exit":
		handleExit()
	default:
		// command not found
		fmt.Println("command not found: ", cmd[0])
	}
	return !exitShell
}
