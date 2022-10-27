package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("*** Linux Interative Shell ***")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayPrompt()
		scanner.Scan()

		// get the input command
		input := scanner.Text()
		input_tokens := strings.Fields(input)
		command, argument := input_tokens[0], input_tokens[1:]

		executeCommand(command, argument)
	}
}
