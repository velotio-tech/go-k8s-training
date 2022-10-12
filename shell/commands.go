package shell

import (
	"fmt"
	"os"
	"strconv"
)

func ls(args ...string) (result string, err error) {
	printDirName := false
	if len(args) == 0 {
		args = append(args, ".")
	}
	if len(args) > 1 {
		printDirName = true
	}
	for _, arg := range args {
		result += fmt.Sprintln()
		if printDirName {
			result += fmt.Sprintln(arg, ":")
		}
		dirEntries, err := os.ReadDir(arg)
		if err != nil {
			return result, err
		}
		for _, dirEntry := range dirEntries {
			result += fmt.Sprintln(dirEntry.Name())
		}
	}
	return
}

func pwd(args ...string) (result string, err error) {
	if len(args) != 0 {
		return "", fmt.Errorf("too many arguments")
	}
	return os.Getwd()
}

func exit(args ...string) (result string, err error) {
	var statusCode int64 = 0
	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments")
	}
	if len(args) == 1 {
		statusCode, _ = strconv.ParseInt(args[0], 10, 64)
	}
	os.Exit(int(statusCode))
	return "", nil
}

func cd(args ...string) (result string, err error) {
	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments")
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if len(args) == 1 {
		dir = args[0]
	}
	err = os.Chdir(dir)
	if err != nil {
		return "", err
	}
	currentDir, _ := pwd()
	return currentDir, nil
}

