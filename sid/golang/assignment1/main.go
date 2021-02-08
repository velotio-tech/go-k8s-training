package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/farkaskid/go_assignments/assignment1/handlers"
	"github.com/farkaskid/go_assignments/assignment1/helpers"
)

func main() {
	cmdReader := bufio.NewReader(os.Stdin)
	hist_size, err := strconv.Atoi(strings.TrimSpace(os.Getenv("HISTSIZE")))

	if err != nil {
		fmt.Println("Failed to get the history size, defaulting to 1000", err)

		hist_size = 1000
	}

	history := &helpers.History{}
	history.Init(int(hist_size))

	for {
		if takeInput(cmdReader, history) {
			break
		}
	}

	fmt.Println("Bye! 👋")
}

func takeInput(cmdReader *bufio.Reader, history *helpers.History) bool {
	// Takes the input from stdin

	username := helpers.GetCurrentUser()
	hostname := helpers.GetHostname()
	cwd := helpers.Getcwd()

	fmt.Printf("%v@%v %v --> ", username, hostname, cwd)

	line, err := cmdReader.ReadString(byte('\n'))

	if err != nil {
		helpers.PrintError("Failed to read console input", err)

		return true
	}

	line = strings.TrimSpace(line)

	history.Record(line)

	if strings.Compare("exit", line) == 0 {
		return true
	}

	handlers.Handle(line, history)

	return false
}
