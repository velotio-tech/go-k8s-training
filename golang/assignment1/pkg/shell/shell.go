package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func executeShell() {
	currentUser, _ := user.Current()
	hostname, _ := os.Hostname()
	scanner := bufio.NewScanner(os.Stdin)
	wd, _ := os.Getwd()

	// print the prefix on console
	fmt.Print(currentUser.Name, "@", hostname, wd)

	//take input
	scanner.Scan()

	//execute the command
	executeCommand(scanner.Text())
}

func executeCommand(s string) {
	cmd := strings.Split(s, " ")
	switch cmd[0] {
	case "ls":
		handleList()
		break
	case "pwd":
		handlePwd()
		break
	case "exit":
		handleExit()
		break
	case "cd":
		handleChangeDir(cmd)
		break
	// case "history":
	// 	handleHistory()
	// 	break
	default:
		fmt.Println(cmd[0], " command not found")
		break
	}
}

// handles the 'ls' command
func handleList() {
	wd, _ := os.Getwd()
	files, _ := os.ReadDir(wd)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// handles the 'pwd' command
func handlePwd() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
}

//handles the 'exit' command
func handleExit() {
	fmt.Println("Exiting Shell...")
	os.Exit(0)
}

//handles the 'cd' command
func handleChangeDir(cmd []string) {
	if len(cmd) < 2 {
		return
	}
	err := os.Chdir(cmd[1])
	if err != nil {
		fmt.Println(err.Error())
	}

}

//hanldes the 'history' command
// func handleHistory() {

// }
