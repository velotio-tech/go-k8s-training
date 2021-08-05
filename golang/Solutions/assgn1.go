package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var file string = "history.txt"

func check(e error) {
	if e != nil {
		os.Truncate(file, 0)
		panic(e)
	}
}

func executeCmd(cmd string) string {
	out, err := exec.Command(cmd).Output()
	check(err)
	return strings.TrimSuffix(string(out), "\n")
}

func addToHistory(cmd string) {
	// write to file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	_, err = f.Write([]byte(cmd + "\n"))
	check(err)
	f.Close()
}

func main() {

	var header string
	var cmd string

	for {
		header = ""
		header += executeCmd("whoami") + "/" + executeCmd("hostname") + executeCmd("pwd") + "~"
		fmt.Print(header + " ")

		reader := bufio.NewReader(os.Stdin)
		read, _ := reader.ReadString('\n')
		cmd = strings.Replace(read, "\n", "", -1)

		if string(cmd[0]) != " " {
			addToHistory(cmd)
		} else {
			cmd = strings.TrimSpace(cmd)
		}

		split := strings.Split(cmd, " ")

		switch split[0] {
		case "exit":
			os.Truncate(file, 0)
			os.Exit(0)
		case "cd":
			os.Chdir(split[1])
		case "history":
			// reaad file here
			data, err := ioutil.ReadFile(file)
			check(err)
			fmt.Println(string(data))
		default:
			out, err := exec.Command(split[0], split[1:]...).Output()
			if err != nil {
				fmt.Println("This command is not supported")
			} else {
				fmt.Println(string(out) + "\n")
			}
		}
	}
}
