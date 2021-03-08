package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

// GetCWD returns current working directory
func GetCWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: can't get current working directory")
		os.Exit(1)
	}
	return cwd
}

// GetHistory loads history of previous sessions from file
func GetHistory() []string {
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

func getHomeDir() string {
	// returns current user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return homeDir
}

// SaveHistoryToFile saves history to file for persisting storage
func SaveHistoryToFile(h []string) {
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

// SaveHistory saves history in memory for each command
func SaveHistory(h []string, text string) []string {
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

// GetPrompt return prompt string user@hostname:pwd$
func GetPrompt() string {
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
	cwd := GetCWD()
	cwd = replaceHomeWithSymbol(cwd)
	return user.Username + "@" + hostname + ":" + cwd + "$ "
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
