/*
this is solution for go lang assignment 1
it supports following commands
1) ls
2) pwd
3) cd
4) history
5) exit

for persisting storage of historic commands it creates file named ".go_assignment_history_file"
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/helpers"
)

func main() {
	// get history "[]string" each elem is a command
	history := helpers.GetHistory()

	// main loop which will wait for next command to be entered by user
	for {
		// get prompt string
		prompt := helpers.GetPrompt()
		fmt.Printf(prompt)
		// read stdin
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// removing last new line char
		text = text[:len(text)-1]
		// get slice of words by splitting on whitespace
		args := strings.Fields(text)
		if len(args) == 0 {
			//nothing is typed so continuing
			continue
		}

		// process command with number (as showed by history command)
		if strings.HasPrefix(args[0], "!") == true {
			// get string excluding first char i.e ! and try if it's a digit
			commandNumber := args[0][1:]
			index, err := strconv.Atoi(commandNumber)
			if err != nil {
				fmt.Println("Error ", err)
			} else {
				// checking if the history exists for the index-1
				i := index - 1
				if i >= 0 && i < len(history) {
					// overwriting args and command text as per the history line
					args = strings.Fields(history[i])
					text = history[i]
					// printing the command first
					fmt.Println(history[i])
				}
			}
		}
		// save history in memory
		history = helpers.SaveHistory(history, text)

		// switch on the command
		switch args[0] {
		case "ls":
			helpers.HandleLS(args)
		case "cd":
			helpers.HandlCD(args)
		case "pwd":
			if len(args) > 1 {
				fmt.Println("Error: too many arguments")
				continue
			}
			cwd := helpers.GetCWD()
			fmt.Println(cwd)
		case "history":
			if len(args) > 1 {
				fmt.Println("Erro: too many arguments")
				continue
			}
			// show history
			for i, h := range history {
				fmt.Printf("%v  %v\n", i+1, h)
			}
		case "exit":
			// upon exit save the history to file
			helpers.SaveHistoryToFile(history)
			os.Exit(0)
		default:
			fmt.Printf("%v: command not found\n", args[0])
		}

	}

}
