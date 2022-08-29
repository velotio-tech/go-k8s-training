package main

import (
	"fmt"
	"os"
	"strings"
	"os/user"
)


func executeCMD(input string, history map[int]string, validCMD map[string]string) bool {

	// split the command and arguments
	fields := strings.Fields(input)
	command := fields[0]

	switch command {
	case "ls":
		executeLS()
	case "pwd":
		executePWD()
	case "history":
		executeHISTORY(history)
	case "cd":
		executeCD(fields)
	case "help":
		executeHELP(validCMD)
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("command not found: '%s'\n", fields[0])
		fmt.Printf("type 'help' to see all available commands\n")
	}

	return true
}

// to handle 'pwd' command
func executePWD() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error while executing PWD command : ",err)
		return
	}
	fmt.Println(pwd)
}

// to handle 'ls' command
func executeLS() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error while executing LS command : ",err)
		return
	}
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// to handle 'cd' command
func executeCD(path []string) {
	if len(path) <=1 {
		fmt.Println("error while executing CD command : Invalid arguments")
		return
	}
	err := os.Chdir(path[1])
	if err != nil {
		fmt.Println("error while executing CD command : ",err)
		return
	}
}

// add command to the history
func addToHistory(command string, history map[int]string) {
	length := len(history)
	history[length+1] = command
}

// to handle 'history' command
func executeHISTORY(historyMap map[int]string) {
	for idx, history := range historyMap {
		fmt.Printf("%d\t%s\n", idx, history)
	}
}

// to handle 'help' command
func executeHELP(validCMD map[string]string) {
	for command, info := range validCMD {
		fmt.Printf("\n%s : %s\n", command, info)
	}
}

// function to generate prefix
func getPrefix() string {
	user, _ := user.Current()
	username := user.Username
	hostname, _ := os.Hostname()
	path, _ := os.Getwd()
	prefix := username + "@" + hostname + ":~" + path + ": "
	return prefix
}