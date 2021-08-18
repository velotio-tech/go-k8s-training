package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

var (
	currUser, _ = user.Current()
	hostname, _ = os.Hostname()
	scanner     = bufio.NewScanner(os.Stdin)
	colorGreen  = "\033[32m"
	colorWhite  = "\033[37m"
)

//	starts the shell
func LaunchShell() {
	fmt.Println("starting shell...")
	os.Chdir("/home/Prasad")
	for {
		printCommandPrefix()
		scanner.Scan()
		execCommand(scanner.Text())
	}
}

//	Executes the shell command.
func execCommand(cmdStr string) {
	cmdStr = strings.Trim(cmdStr, " ")
	tokens := strings.Split(cmdStr, " ")

	switch tokens[0] {
	case "ls":
		handleListFilesCommand()
		break

	case "cd":
		handleChangeDirCommand(tokens)
		break

	case "pwd":
		handlePWDCommand()
		break

	case "exit":
		fmt.Println("exiting shell...")
		os.Exit(0)
		break

	default:
		fmt.Println("Need to handle invalid commands.")
		break
	}
}

//	handles the 'ls' command of shell.
func handleListFilesCommand() {
	wd, _ := os.Getwd()
	files, _ := os.ReadDir(wd)

	for _, file := range files {
		fileInfo, _ := file.Info()
		fmt.Println(fileInfo.Size(), " ", file.Name())
	}
}

//	handles the 'cd' command of shell
func handleChangeDirCommand(tokens []string) {
	if len(tokens) < 2 {
		return
	}
	err := os.Chdir(tokens[1])
	if err != nil {
		fmt.Println(err.Error())
	}
}

//	handles the 'pwd' command
func handlePWDCommand() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
}

//	Prints the command prefix on consle.
func printCommandPrefix() {
	wd, _ := os.Getwd()
	fmt.Print(string(colorGreen), currUser.Username, "@", hostname, ":", wd, "$ ")
	fmt.Print(string(colorWhite), "")
}
