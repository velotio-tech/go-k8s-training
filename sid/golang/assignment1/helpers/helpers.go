package helpers

import (
	"fmt"
	"os"
	"os/user"
)

// Red color for error outs
const colorRed = "\033[31m"
const colorBase = "\033[0m"

func PrintError(msg string, err error) {
	// Prints the error in Red color

	fmt.Println(string(colorRed), msg+": "+err.Error(), string(colorBase))
}

func GetCurrentUser() string {
	// Gets current logged in user

	user, err := user.Current()

	if err != nil {
		PrintError("Failed to get the current user information", err)

		return ""
	}

	return user.Username
}

func GetHostname() string {
	// Gets the hostname

	hostname, err := os.Hostname()

	if err != nil {
		PrintError("Failed to get the hostname", err)

		return ""
	}

	return hostname
}

func Getcwd() string {
	// Gets the current working directory

	cwd, err := os.Getwd()

	if err != nil {
		PrintError("Failed to get the current working directory", err)

		return ""
	}

	return cwd
}
