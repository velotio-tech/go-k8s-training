package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
)

type history struct {
	command *list.List
	size    int
}

func (h *history) init(size int) {
	h.command = list.New()
	h.size = size
}

func (h *history) capture(cmd string) {
	h.command.PushFront(cmd)
	if h.command.Len() > h.size {
		h.command.Remove(h.command.Back())
	}
}

func getTerminal() string {

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
	terminal := fmt.Sprintf("%s@%s:~/%s", user.Username, hostName, strings.Join(cwd[3:], "/"))
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
	return getTerminal()
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

func getHistory(h *history) string {
	output := ""
	for item := h.command.Front(); item != nil; item = item.Next() {
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

func main() {
	fmt.Println(getTerminal())

	histSize, err := strconv.Atoi(strings.TrimSpace(os.Getenv("HISTSIZE")))

	if err != nil {
		histSize = 1000
	}

	history := history{}
	history.init(int(histSize))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		history.capture(line)

		if strings.Compare("exit", line) == 0 {
			os.Exit(0)
		}

		cmd := strings.Fields(line)

		output := ""

		switch cmd[0] {

		case "cd":
			output = getCd(cmd, line)

		case "ls":
			output = getLs()

		case "pwd":
			output = getPwd()

		case "history":
			output = getHistory(&history)

		default:
			output = fmt.Sprintf("Command %q not found", line)
		}

		fmt.Println(output)
	}

}
