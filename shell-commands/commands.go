package main

import (
	"fmt"
	"os"
)

// To get current working directory
func getWorkingDirectory() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)
}

// To handle ls command
func getListing() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// To change directory
func changeDirectory(path []string) {
	if len(path) <=1 {
		fmt.Println("Invalid arguments to command")
		return
	}
	err := os.Chdir(path[1])
	if err != nil {
		fmt.Println(err)
	}
}

// Adds entry to historyMap after every command hit
func addHistory(command string) {
	if historyMap == nil {
		historyMap = make(map[int]string)
	}
	length := len(historyMap)
	historyMap[length+1] = command
}

// To get history
func getHistory() {
	for idx, history := range historyMap {
		fmt.Printf("%d\t%s\n", idx, history)
	}
}
