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

	fmt.Println("------------------Welcome to Linux Shell--------------------")

	scanner := bufio.NewScanner(os.Stdin)

	historyObj := restoreHistory()

	for {

		fmt.Print(getPromt())
		scanner.Scan()
		inputCommand := scanner.Text()
		tokens := strings.Fields(inputCommand)
		command, argument := tokens[0], tokens[1:]

		switch command {

		case "exit":
			historyObj.update(command)
			historyObj.saveHistoryToFile(command)
			fmt.Println("exit for shell")
			os.Exit(0)
		case "pwd":
			historyObj.update(command)

			fmt.Println(presentWorkingDir())
		case "ls":
			historyObj.update(command)

			listDirectoriesAndFiles()
		case "cd":
			historyObj.update(command + " " + argument[1])
			changeDir(argument)
		case "history":
			historyObj.update(command)

			historyObj.display()
		default:
			historyObj.update(command)
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

func getPromt() string {

	currUser, _ := user.Current()
	hostname, _ := os.Hostname()
	str := currUser.Username + "@" + hostname + ":~"

	path := presentWorkingDir()

	arr := strings.Split(path, "/")

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
