package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func getPrompt() string {
	currentUser, _ := user.Current()
	hostname, _ := os.Hostname()

	return "[" + currentUser.Username + "@" + hostname + "]:" + getCurrentDirectory() + " $ "
}

func getCurrentDirectory() string {
	directories := strings.Split(presentWorkingDirectory(), "/")

	return directories[len(directories)-1]
}

func printFilesAndDirectories() {
	files, _ := os.ReadDir(".")

	if len(files) != 0 {
		for _, file := range files {
			fmt.Print(file.Name(), "\t\t")
		}
		fmt.Println()
	}
}

func presentWorkingDirectory() string {
	currentDirectoryPath, _ := os.Getwd()
	return currentDirectoryPath
}
