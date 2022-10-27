package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
)

var file string = "history.txt"
var i int = 0

func displayPrompt() {
	// get the current user
	currentUser, _ := user.Current()

	// get the hostname
	hostname, _ := os.Hostname()

	// get the current working directory
	pwd := showPresentWorkingDir()

	prompt := currentUser.Username + "@" + hostname + " ~" + pwd + "$ "
	fmt.Print("\n" + prompt)
}

func executeCommand(cmd string, args []string) {
	fmt.Print("\nExecuting command :", cmd+"\n")
	fmt.Print("-----------------------\n")

	// add the command into file
	i++
	updateHistoryFile(strconv.Itoa(i) + " " + cmd)

	if len(cmd) == 0 {
		fmt.Println("Error : not received any command, exit!")
	}

	switch cmd {
	case "ls":
		showFilesAndDirectories(args)

	case "cd":
		processChangeWorkingDir(args)

	case "pwd":
		fmt.Println(cmd, ": Current working directory : ", showPresentWorkingDir())

	case "history":
		// read file here
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error : Unable to read a file!")
		}
		fmt.Println(string(data))

	case "exit":
		fmt.Println("\nExiting from the shell. Thank you!")
		os.Remove(file)
		os.Exit(0)

	default:
		fmt.Println(cmd, " command not found..!")
	}
}

// Display current working directory
func showPresentWorkingDir() string {

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error : ", err)
	}

	return dir
}

// Change the working directory
func processChangeWorkingDir(args []string) {
	if len(args) == 0 {
		fmt.Println("Error : Invalid command argument!")
	}

	err := os.Chdir(args[0])

	if err != nil {
		fmt.Println("Error : Directory not found")
	}
}

// List the files and directories
func showFilesAndDirectories(args []string) {

	// command -> ls
	if len(args) == 0 {
		files, err := ioutil.ReadDir(showPresentWorkingDir())
		if err != nil {
			fmt.Println("Error : No file or directory!", err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}

	// command -> ls ~/path
	for _, path := range args {

		fmt.Println("Path :", path)
		_, err := os.Stat(path)

		if err != nil {
			fmt.Println("Error : no path found!")
			return
		}

		file, _ := os.ReadDir(path)
		for _, file := range file {
			fmt.Println(file.Name())
		}
	}
}

// Add commands to the history file
func updateHistoryFile(command string) {
	// open and write into a file
	file, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error : open operation failed!")
	}

	_, err = file.Write([]byte(command + "\n"))
	if err != nil {
		fmt.Println("Error : ", err)
	}
	file.Close()
}
