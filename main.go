package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {

	fmt.Println("-------------------------------Welcome to Linux Shell-------------------------------------------\n")

	var inputCommand string

	for {
		fmt.Print(getPromt())
		fmt.Scanln(&inputCommand)

		inputCommand = strings.ToLower(inputCommand)

		if inputCommand == "ls" {
			listDirectoriesAndFiles()
		} else if inputCommand == "exit" {
			os.Exit(0)
		} else if inputCommand == "pwd" {
			presentWorkingDir()
		}
	}

}

func presentWorkingDir() {
	currentDirPath, _ := os.Getwd()

	fmt.Print(currentDirPath)

}

func listDirectoriesAndFiles() {
	files, _ := os.ReadDir(".")

	for _, file := range files {
		fmt.Print(file.Name(), "\t")
	}
}

func getPromt() string {
	currUser, _ := user.Current()

	return "[" + currUser.Username + " ] " + getDirectory() + " $ "
}

func getDirectory() string {
	currentDirPath, _ := os.Getwd()

	directories := strings.Split(currentDirPath, "/")

	return directories[len(directories)-1]
}
