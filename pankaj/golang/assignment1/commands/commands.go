package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		return getLs()

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
		log.Fatal("Failed to get the current user information.", err)

	}
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal("Failed to get host name.", err)

	}
	getWd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get the current working directory.", err)
	}
	cwd := strings.Split(getWd, "/")
	terminal := fmt.Sprintf("%s@%s:~/%s ", user.Username, hostName, strings.Join(cwd[3:], "/"))
	return terminal
}

func getCd(cmd []string, line string) string {
	dst := cmd[1]

	if strings.HasPrefix(dst, "~") {
		dst = strings.Replace(dst, "~", os.Getenv("HOME"), 1)
	}

	err := os.Chdir(dst)

	if err != nil {
		return fmt.Sprintf("Command %q not found", line)

	}
	return ""
}

func getLs() string {
	dirs, err := ioutil.ReadDir("./")
	if err != nil {
		return fmt.Sprintf("Failed to get current working dirs %v.", err)
	}
	output := ""
	for _, file := range dirs {
		output = output + " " + file.Name()
	}
	return output
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
