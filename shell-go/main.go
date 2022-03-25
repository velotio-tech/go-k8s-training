package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func main() {

	// -----------Variable declaration part starts-----------------
	
	// Create a map to save history of commands executed
	// It'll store the count (int) & the command (string)
	m := make(map[int]string)

	//counter variable to keep count of commands
	counter := 0

	// decalare color variables to distingish between pretext and actaul command
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"

	// to read input from standard input
	reader := bufio.NewReader(os.Stdin)

	// ------------Variable decelration part ends-------------------


	//Run an infinite loop till it encounters an exit command
	for {

		//increment the counter
		counter++

		// pretext consists of username + hostname + current working directory
		pretext := preText()

		// use yellow color for pretext
		fmt.Print(string(colorYellow), pretext)

		// use blue color for commands
		fmt.Print(string(colorBlue))

		// read until new line is encounterd (new line is added with Enter key)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// store the command in the map
		m[counter] = input

		// Execute the entered command
		if err = ExecInput(input, m, counter); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}


//function to generate pretext
// pretext consists of username + hostname + current working directory
func preText() string {
	user, _ := user.Current()
	username := user.Username
	hostname, _ := os.Hostname()
	path, _ := os.Getwd()
	pretext := username + "@" + hostname + ":~" + path + " "
	return pretext
}
