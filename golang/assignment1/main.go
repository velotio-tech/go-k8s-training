
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var history []string
var shell string

func main() {
	currentUser := getCurrentUser()
	hostName := getHostName()
	currDir := getCurrDir()

	for {
		shell = currentUser + "@" + hostName + " " + currDir + " % "
		fmt.Print(shell)
		reader := bufio.NewReader(os.Stdin)
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		if cmdStr != "" && cmdStr != "\n" {
			history = append(history, cmdStr)
		}
		runCmd(cmdStr)
	}
}

func getCurrentUser() string {
	currUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error while getting current user", err)
	}
	return currUser.Username
}

func getHostName() string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error while getting hostname ", err)
	}
	return host
}

func getCurrDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

func changeDir(params []string) {
	var dir = ""
	if len(params) == 1 {
		dir = "/Users"
	} else {
		dir = params[1]
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Print(shell)
		fmt.Println(err)
	} else {
		fmt.Print(shell)
		fmt.Println(getCurrDir())
	}
}

func getHistory() {
	for i, cmd := range history {
		fmt.Println(i+1, cmd)
	}
}

func runCmd(cmdStr string) {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	params := strings.Fields(cmdStr)
	switch params[0] {
	case "exit":
		os.Exit(0)
	case "history":
		getHistory()
	case "cd":
		changeDir(params)
	default:
		cmd := exec.Command(cmdStr)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(shell)
		fmt.Print(outb.String())
	}

}

