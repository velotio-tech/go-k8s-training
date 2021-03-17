package commands

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment1/helpers"
)

func CommandHandler(command string, history *helpers.History) string {
	cmd := strings.Fields(command)
	switch cmd[0] {

	case "cd":
		return getCd(cmd, command)

	case "ls":
		return getLs(cmd[1:])

	case "pwd":
		return getPwd()

	case "history":
		return getHistory(history)

	default:
		return fmt.Sprintf("Command %q not found", command)
	}

}
func GetTerminal() string {

	user, err := user.Current()
	if err != nil {
		fmt.Println("Failed to get the current user information.", err.Error())

	}
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get host name.", err.Error())

	}
	getWd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the current working directory.", err.Error())
	}
	terminal := fmt.Sprintf("%s@%s:~%s$ ", user.Username, hostName, getWd)
	return terminal
}

func getCd(cmd []string, line string) string {
	dst := ""
	if len(cmd) == 1 {
		dst = os.Getenv("HOME")
	} else if len(cmd) == 2 {
		dst = cmd[1]
		if strings.HasPrefix(dst, "~") {
			dst = strings.Replace(dst, "~", os.Getenv("HOME"), 1)
		}
	} else {
		fmt.Println("cd: too many arguments")
		return ""
	}

	err := os.Chdir(dst)

	if err != nil {
		fmt.Println("Failed to change working directory", err.Error())

	}
	return ""
}

func getLs(command []string) string {
	var cmd []string
	if len(command) == 0 {
		cmd = append(cmd, "./")
	} else {
		cmd = append(cmd, command...)
	}
	dirs, err := exec.Command("ls", cmd...).Output()
	if err != nil {
		return fmt.Sprintf("Failed to get current working dirs %v.", err)
	}
	return string(dirs[:])
}

func getHistory(h *helpers.History) string {
	output := ""
	for item := h.Command.Front(); item != nil; item = item.Next() {
		output = output + "\n" + fmt.Sprintf("%v", item.Value)
	}
	return output
}

func getPwd() string {
	output, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("Failed to get the current working directory %v.", err)
	}
	return output
}
