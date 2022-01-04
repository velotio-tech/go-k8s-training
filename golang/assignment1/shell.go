package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

func executeShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		dailogBox()
		//take input from Stdin
		scanner.Scan()
		// execute the command
		isExit := executeCommand(scanner.Text())
		if isExit {
			break
		}
	}
}
func dailogBox() {
	currentUser, _ := user.Current()
	pwd, _ := os.Getwd()
	hostName, _ := os.Hostname()
	fmt.Printf("%v@%v %v ", currentUser.Username, hostName, pwd)
}

func executeCommand(s string) bool {
	const shellExit = true
	cmd := strings.Fields(s)
	if len(cmd) == 0 {
		return !shellExit
	}
	switch cmd[0] {
	case "ls":
		handleList(cmd[1:])
	case "pwd":
		handlePwd()
	case "exit":
		handleExit()
	case "cd":
		handleChangeDir(cmd[1:])
	default:
		fmt.Println(cmd[0], " command not found")
	}
	return !shellExit
}

// handles the 'ls' command
func handleList(cmd []string) {
	if len(cmd) == 0 {
		wd, _ := os.Getwd()
		files, err := os.ReadDir(wd)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
	for _, path := range cmd {
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

// handles the 'pwd' command
func handlePwd() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wd)
}

//handles the 'exit' command
func handleExit() {
	fmt.Println("Exiting Shell...")
	os.Exit(0)
}

//handles the 'cd' command
func handleChangeDir(cmd []string) {
	if len(cmd) != 1 {
		fmt.Println("Invalid number of arguments")
	}
	err := os.Chdir(cmd[0])
	if err != nil {
		fmt.Println(err)
	}

}
