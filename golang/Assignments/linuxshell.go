package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func perform(actionString string, history *[]string) {
	action := strings.Split(actionString, " ")
	switch action[0] {
	case "ls":
		if len(action) < 2 {
			out, err := os.ReadDir(".")
			for _, file := range out {
				fmt.Println(file.Name())
			}
			if err != nil {
				fmt.Printf("%s", err)
			}

		} else {
			out, err := os.ReadDir(action[1])
			if err != nil {
				fmt.Printf("%s", err)
			}
			for _, file := range out {
				fmt.Println(file.Name())
			}
		}

		*history = append(*history, actionString)
	case "pwd":
		out, err := os.Getwd()
		if err != nil {
			fmt.Printf("%s", err)
		}
		output := string(out[:])
		fmt.Println(output)
		*history = append(*history, "pwd")
	case "cd":
		os.Chdir(action[1])
		*history = append(*history, actionString)
	case "history":
		for _, cmd := range *history {
			fmt.Println(cmd)
		}
		*history = append(*history, "history")
	case "exit":
		fmt.Printf("Exiting Program..\n")
		os.Exit(1)
	default:
		*history = append(*history, actionString)
		fmt.Printf("Invalid Command: %v\n Try Again!!\n", action)
	}
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can not execute linux command on windows platform")

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var cmdlist []string
		user, _ := user.Current()
		hostname, _ := os.Hostname()
		for {
			cwd, _ := os.Getwd()
			fmt.Printf("%v@%v %v: ", user.Username, hostname, cwd)
			scanner.Scan()
			perform(scanner.Text(), &cmdlist)
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}

	}
}
