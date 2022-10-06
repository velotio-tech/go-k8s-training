package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

type history []string

func main() {

	history := history{}

	for {
		shell := getCurrentUser() + "@" + getHostName() + " " + getCurrentDir() + " % "
		fmt.Print(shell)
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		history = append(history, command)
		runCmd(command, history)
	}
}

func runCmd(command string, h history) {
	command = strings.TrimSuffix(command, "\n")
	params := strings.Fields(command)

	switch params[0] {
	case "exit":
		os.Exit(0)
	case "ls":
		listDir()
	case "pwd":
		fmt.Println(getCurrentDir())
	case "cd":
		changeDir(params)
	case "history":
		fmt.Println(h)
	default:
		fmt.Println("Command not found:", params[0])
	}

}

// func newHistory() history {
// 	return history{}
// }

// func printHistory() {
// 	for i, cmd := range h {
// 		fmt.Println(i+1, cmd)
// 	}
// }

func getCurrentUser() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get currect user")
	}
	return user.Username
}

func getHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Unable to get host name")
	}
	return hostName
}

func getCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to get current dir")
	}
	return currentDir
}

func changeDir(params []string) {
	fmt.Println(params)
	err := os.Chdir(params[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getCurrentDir())
	}
}

func listDir() {
	files, err := ioutil.ReadDir(getCurrentDir())
	if err != nil {
		fmt.Println("Unable to list directories", err)
	}
	for _, f := range files {
		deets := fmt.Sprintf("%v, %v", f.Name(), f.Mode())
		fmt.Println(deets)
	}
}
