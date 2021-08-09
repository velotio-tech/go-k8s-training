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
		fmt.Println(e)
	}
}

func splitCmd(cmd string) []string {
	var split []string
	if strings.Contains(cmd, "\"") {
		ind := strings.Index(cmd, "\"")
		split = strings.Fields(cmd[:ind])
		split = append(split, cmd[ind:])
		split[len(split)-1] = strings.Trim(split[len(split)-1], "\"")
	} else {
		split = strings.Split(cmd, " ")
	}
	return split
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
	var split []string

	for {
		header = ""
		header += executeCmd("whoami") + "*" + executeCmd("hostname") + "*" + executeCmd("pwd") + "~"
		fmt.Print(header + " ")

		reader := bufio.NewReader(os.Stdin)
		read, _ := reader.ReadString('\n')
		cmd = strings.Replace(read, "\n", "", -1)

		if cmd == "" {
			continue
		}

		if string(cmd[0]) != " " {
			addToHistory(cmd)
		} else {
			cmd = strings.TrimSpace(cmd)
		}

		split = splitCmd(cmd)

		switch split[0] {
		case "exit":
			os.Truncate(file, 0)
			os.Exit(0)
		case "cd":
			var loc string
			var err error
			if len(split) < 2 {
				loc, err = os.UserHomeDir()
				check(err)
			} else {
				loc = strings.Join(split[1:], "")
				if strings.ContainsAny(loc, "\\") {
					ind := strings.Index(loc, "\\")
					fmt.Println("MM", ind)
					loc = strings.Replace(loc, "\\", " ", -1)
				}
			}
			os.Chdir(loc)
		case "history":
			// read file here
			data, err := ioutil.ReadFile(file)
			check(err)
			fmt.Println(string(data))
		default:
			out, err := exec.Command(split[0], split[1:]...).Output()
			if err != nil {
				fmt.Println("This command is not supported")
				fmt.Println(err)
			} else {
				fmt.Println(string(out) + "\n")
			}
		}
	}
}
