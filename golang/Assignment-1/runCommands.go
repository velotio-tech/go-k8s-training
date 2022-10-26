package main

import (
	"fmt"
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
		response = append(response, file.Name())
	}
	return response
}

func changeDirectory(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New Current Directory : " + getCurrentDirectory())
	}
}

func exitShell() {
	fmt.Println("Exiting the shell...")
	os.Exit(0)
}
