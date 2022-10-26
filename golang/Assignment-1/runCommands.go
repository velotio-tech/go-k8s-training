package main

import (
	"os"
	"os/user"
)

func getPrompt() string {
	currentUser, _ := user.Current()
	hostname, _ := os.Hostname()
	currentDirectory := getCurrentDirectory()

	return currentUser.Username + "@" + hostname + " ~" + currentDirectory
}

func getCurrentDirectory() string {
	currentDirectory, _ := os.Getwd()
	return currentDirectory
}

func getListOfFilesAndDirectories() []string {
	files, _ := os.ReadDir(".")
	var response []string
	for _, file := range files {
		// fmt.Print(file.Name(), "\t")
		response = append(response, file.Name())
	}
	return response
}

func exitShell() {
	os.Exit(0)
}
