package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment1/commands"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment1/helpers"
)

func main() {
	terminal := commands.GetTerminal()
	fmt.Printf(terminal)

	histSize, err := strconv.Atoi(strings.TrimSpace(os.Getenv("HISTSIZE")))

	if err != nil {
		histSize = 1000
	}

	history := helpers.History{}
	history.Init(int(histSize))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		history.Capture(line)

		if strings.Compare("exit", line) == 0 {
			os.Exit(0)
		}

		output := commands.CommandHandler(line, &history)

		fmt.Println(output)
		fmt.Printf(commands.GetTerminal())
	}

}
