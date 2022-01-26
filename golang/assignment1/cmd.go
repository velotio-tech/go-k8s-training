package main

import (
	"fmt"
	"log"
	"os"
)

// handling the 'ls' linux command
func handleLs(args []string) {
	if len(args) == 0 {
		workDir, _ := os.Getwd()
		files, err := os.ReadDir(workDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
	for _, path := range args {
		_, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		} else {
			files, err := os.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
	}
}

// handling the 'pwd' linux command
func handlePwd() {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(workDir)
}

// handling the 'exit' linux command
func handleExit() {
	fmt.Println("Exiting Shell")
	os.Exit(0)
}

// handling the 'cd' linux command
func handleCd(args []string) {
	if len(args) != 1 {
		fmt.Println("Invalid arguments")
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println(err)
	}

}
