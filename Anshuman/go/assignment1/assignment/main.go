package main

import (
	"assignment1/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var splitString []string

	for {
		fmt.Printf("%s", utils.GetSignature())
		text, _ := reader.ReadString('\n')
		utils.RecordHistory(text)
		text = strings.ReplaceAll(text, "\n", "")
		text = strings.Trim(text, " ")
		splitString = strings.Split(text, " ")
		cmdLen := len(splitString)
		switch splitString[0] {
		case "pwd":
			if cmdLen != 1 {
				fmt.Printf(utils.ErrorSt, "pwd")
			} else {
				utils.PWD()
			}
		case "cd":
			if cmdLen == 1 {
				utils.CD("")
			} else if cmdLen == 2 {
				utils.CD(splitString[1])
			} else {
				fmt.Printf(utils.ErrorSt, "cd")
			}
		case "history":
			if cmdLen != 1 {
				fmt.Printf(utils.ErrorSt, "history")
			} else {
				utils.History()
			}
		case "ls":
			if cmdLen == 1 {
				utils.LS(".")
			} else if cmdLen == 2 {
				utils.LS(splitString[1])
			} else {
				fmt.Printf(utils.ErrorSt, "ls")
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("invalid command entered")
			utils.ShowValidCommands()
		}
	}
}
