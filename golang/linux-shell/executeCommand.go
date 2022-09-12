package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func getPrompt() string {
	currentUser, _ := user.Current()

	return "[" + currentUser.Username + "]:" + getCurrentDirectory() + " $ "
}

func getCurrentDirectory() string {
	currentDirPath, _ := os.Getwd()
	directories := strings.Split(currentDirPath, "/")

	return directories[len(directories)-1]
}

func listDirectoriesAndFiles() {
	files, _ := os.ReadDir(".")

	for _, file := range files {
		fmt.Print(file.Name(), "\t")
	}
}
