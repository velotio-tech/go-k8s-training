package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"errors"
	"os/user"
)

func main() {

	pretext := getPretext()
		
	for {
		fmt.Print(pretext)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		inputError := runCommand(input)
		
		if inputError != nil {
			fmt.Println(inputError)
		}
	}
}

func getPretext() string {
	user, _ := user.Current()
	username := user.Username
	hostname, _ := os.Hostname()
	path, _ := os.Getwd()
	pretext := username + "@" + hostname + ":~" + path + " "

	return pretext
}

func runCommand(inputString string)  error {
	inputString = strings.TrimSuffix(inputString, "\n")
	
	args := strings.Fields(inputString);

	switch args[0] {
	case "cd": 
		if(len(args) < 2) {
			return errors.New("path required")
		}

		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}
