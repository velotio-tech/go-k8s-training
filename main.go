package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {

	fmt.Println("-------------------------------Welcome to Linux Shell-------------------------------------------\n")

	//var inputCommands string

	//for {

	fmt.Println(getPromt())

	//}

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
