/*
this is solution for go lang assignment 1
it supports following commands
1) ls
2) pwd
3) cd
4) history
5) exit

for persisting storage of historic commands it creates file named ".go_assignment_history_file"
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	// get history "[]string" each elem is a command
	history := getHistory()

	// main loop which will wait for next command to be entered by user
	for {
		// get prompt string
		prompt := getPrompt()
		fmt.Printf(prompt)
		// read stdin
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// removing last new line char
		text = text[:len(text)-1]
		// get slice of words by splitting on whitespace
		args := strings.Fields(text)
		if len(args) == 0 {
			//nothing is typed so continuing
			continue
		}

		// process command with number (as showed by history command)
		if strings.HasPrefix(args[0], "!") == true {
			// get string excluding first char i.e ! and try if it's a digit
			commandNumber := args[0][1:]
			index, err := strconv.Atoi(commandNumber)
			if err != nil {
				fmt.Println("Error ", err)
			} else {
				// checking if the history exists for the index-1
				i := index - 1
				if i >= 0 && i < len(history) {
					// overwriting args and command text as per the history line
					args = strings.Fields(history[i])
					text = history[i]
					// printing the command first
					fmt.Println(history[i])
				}
			}
		}
		// save history in memory
		history = saveHistory(history, text)

		// switch on the command
		switch args[0] {
		case "ls":
			handleLS(args)
		case "cd":
			handlCD(args)
		case "pwd":
			if len(args) > 1 {
				fmt.Println("Error: too many arguments")
				continue
			}
			cwd := getCWD()
			fmt.Println(cwd)
		case "history":
			if len(args) > 1 {
				fmt.Println("Erro: too many arguments")
				continue
			}
			// show history
			for i, h := range history {
				fmt.Printf("%v  %v\n", i+1, h)
			}
		case "exit":
			// upon exit save the history to file
			saveHistoryToFile(history)
			os.Exit(0)
		default:
			fmt.Printf("%v: command not found\n", args[0])
		}

	}

}

func saveHistoryToFile(h []string) {
	// saves history to file for persisting storage
	homeDir := getHomeDir()
	historyFilePath := homeDir + "/.go_assignment_history_file"
	// create or overrite the file
	file, err := os.Create(historyFilePath)
	if err != nil {
		fmt.Println("Error: creating history file", err)
	}
	// join every command with new line char for storing in file
	file.WriteString(strings.Join(h, "\n"))
	// finally close the file
	file.Close()

}
func getHistory() []string {
	// loads history of previous sessions from file
	history := make([]string, 0)
	homeDir := getHomeDir()
	historyFilePath := homeDir + "/.go_assignment_history_file"
	data, err := ioutil.ReadFile(historyFilePath)
	if err != nil {
		// returning empty history slice
		return history
	}
	// splitting with new line char
	return strings.Split(string(data), "\n")
}

func saveHistory(h []string, text string) []string {
	// saves history in memory for each command
	// maximum 50 historic commands will be remembered
	maxLen := 50
	lenOfHistory := len(h)
	// if the current cmd is same as last command don't need to append
	if lenOfHistory > 0 && h[lenOfHistory-1] == text {
		return h
	}
	h = append(h, text)
	// if overflows make it of appropriate size with most recent commands
	if len(h) > maxLen {
		h = h[1:]
	}
	return h

}

func handleLS(args []string) {
	// dirname wil be showed if multiple paths are given to ls
	showDirName := false
	if len(args) == 1 {
		// directory is not given, taking current working directory as default
		cwd := getCWD()
		args = append(args, cwd)
	} else if len(args) > 2 {
		showDirName = true
	}

	for _, path := range args[1:] {
		path = replaceSymbolWithHomeFolder(path)
		if showDirName == true {
			fmt.Println(path + ":\n")
		}

		fs, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("Error: ", err)
			// if error occures for one path try next path
			continue
		}
		for _, each := range fs {
			resetColor := ""
			blueColor := ""
			if each.IsDir() == true {
				resetColor = "\033[0m"
				blueColor = "\033[34m"
			}
			// directory will be showed in blue color
			fmt.Println(blueColor + each.Name() + resetColor)

		}
	}
}

func handlCD(args []string) {
	isValid := true
	if len(args) > 2 {
		fmt.Println("Error: too many arguments")
		isValid = false
	} else if len(args) == 1 {
		// user didn't provide "path" to cd into, default is home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		args = append(args, homeDir)
	}
	if isValid == true {
		// if user uses ~ in path replace with actual home path
		path := replaceSymbolWithHomeFolder(args[1])
		err := os.Chdir(path)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

func getPrompt() string {
	// return prompt string user@hostname:pwd$
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error: can't get current user")
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error: can't get hostname")
		os.Exit(1)
	}
	cwd := getCWD()
	cwd = replaceHomeWithSymbol(cwd)
	return user.Username + "@" + hostname + ":" + cwd + "$ "
}

func getCWD() string {
	// returns current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: can't get current working directory")
		os.Exit(1)
	}
	return cwd
}

func getHomeDir() string {
	// returns current user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return homeDir
}

func replaceHomeWithSymbol(path string) string {
	// replace home dir path with symbol "~"
	homeDir := getHomeDir()
	if strings.HasPrefix(path, homeDir) {
		return strings.Replace(path, homeDir, "~", 1)

	}
	return path
}

func replaceSymbolWithHomeFolder(path string) string {
	// replace symbol "~" with user's home directory
	homeDir := getHomeDir()
	if strings.HasPrefix(path, "~") {
		return strings.Replace(path, "~", homeDir, 1)

	}
	return path
}
