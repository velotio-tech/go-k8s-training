package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecInput(input string, m map[int]string, counter int) error {

	// Trim the "\n" (new line) special charater which get appeneded when we press Enter
	input = strings.TrimSuffix(input, "\n")

	// Split the commands with " " (space) seperator
	// This is done to execute the commands with their paramters. eg: ls -lsd
	args := strings.Split(input, " ")

	// Prepare the command to execute
	// Command returns the Cmd struct to execute the named program with the given arguments.
	cmd := exec.Command(args[0], args[1:]...)

	// handling of commands which are not handled by os/exec package
	switch args[0] {

	//exit from program when "exit" command is executed
	case "exit":
		os.Exit(0)
	
	// mimic "cd" command with Chdir
	// Chdir changes the current working directory to the named directory
	case "cd":
		return os.Chdir(args[1])

	// 
	case "history":
		for i := 1; i <= len(m); i++ {
			fmt.Println(i, m[i])

		}
		return nil
	}

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Run starts the specified command and waits for it to complete.
	return cmd.Run()
}