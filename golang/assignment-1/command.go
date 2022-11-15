package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	ctype string
	args  string
}

func commandHandler(osArgs []string) {

	if len(osArgs) == 1 {
		osArgs = append(osArgs, "")
	}

	args := Command{
		ctype: osArgs[0],
		args:  osArgs[1],
	}
	switch args.ctype {
	case "ls":
		args.ls()
	case "pwd":
		args.pwd()
	case "exit":
		args.exit()
	case "cd":
		args.cd()
	case "history":
		args.history()
	case "clear":
		args.clear()
	default:
		fmt.Println("command not found")
	}

	save(strings.Join(osArgs, " ") + "\n")
}

func (command Command) history() {
	file, err := os.ReadFile("history.log")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s", string(file))
}

func (command Command) ls() {

	path := session.cwd + "/" + command.args

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, file := range files {
		fmt.Printf("%s  ", file.Name())
	}
	fmt.Println()
}

func (command Command) pwd() {
	fmt.Println(session.cwd)
}

func (command Command) exit() {
	fmt.Println("bye!")
	os.Exit(0)
}

func (command Command) cd() {

	var directory string
	if command.args == "" {
		directory, _ = os.UserHomeDir()
	} else {
		directory = session.cwd + "/" + command.args
	}

	err := os.Chdir(directory)
	if err != nil {
		fmt.Println(err)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	session.cwd = dir
}

func (command Command) clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
