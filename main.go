package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func main() {

	fmt.Println("-----------------Welcome to Linux Shell-------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print(getPromt())
		scanner.Scan()
		inputCommand := scanner.Text()
		tokens := strings.Fields(inputCommand)
		command, argument := tokens[0], tokens[1:]

		switch command {

		case "exit":
			fmt.Println("exit for shell")
			os.Exit(0)
		case "pwd":
			fmt.Println(presentWorkingDir())
		case "ls":
			listDirectoriesAndFiles()
		case "cd":
			changeDir(argument)
		default:
			fmt.Println("Command not found:", command)

		}

	}

}

func changeDir(argument []string) {
	os.Chdir(argument[0])
}

func listDirectoriesAndFiles() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	for index, f := range files {
		fmt.Println(index+1, f.Name()+"\n")
	}
}

func presentWorkingDir() string {
	currentDirPath, _ := os.Getwd()
	return currentDirPath
}

func getCurrentDirectory() string {
	currentDirPath, _ := os.Getwd()

	directories := strings.Split(currentDirPath, "/")

	return directories[len(directories)-1]
}

func getPromt() string {

	currUser, _ := user.Current()
	hostname, _ := os.Hostname()
	str := currUser.Username + "@" + hostname + ":~"

	path := presentWorkingDir()

	arr := strings.Split(path, "/")
	//fmt.Println(arr)
	flag := false

	for _, val := range arr {
		if flag {
			str = str + "/" + val
		}

		if val == currUser.Username {
			flag = true
		}

	}

	str = str + "$ "

	return str
}
