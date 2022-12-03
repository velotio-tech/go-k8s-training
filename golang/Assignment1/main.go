package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// Commands
const (
	ls      string = "ls"
	pwd            = "pwd"
	cd             = "cd"
	exit           = "exit"
	help           = "help"
	history        = "history"
)

func getHomedir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home, err = os.Getwd()
		if err != nil {
			panic(err.Error())
		}
	}
	return home
}

var HOME = getHomedir()

// History Filename
const (
	HISTORY_FILE = "/.gohistory"
)

// Create file that saves history command during the startup
func init() {
	createHistoryFile()
}

func createHistoryFile() {
	_, err := os.Create(HOME + HISTORY_FILE)
	if err != nil {
		fmt.Println("Cannot create history file", err.Error())
	}
}

func readHistory() (string, error) {
	history, err := ioutil.ReadFile(HOME + HISTORY_FILE)
	if err != nil {
		return "", err
	}
	return string(history), nil
}

func updateHistory(cmd string) {
	f, err := os.OpenFile(HOME+HISTORY_FILE, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Can't update History", err.Error())
		return
	}
	defer f.Close()
	if _, err := f.WriteString(cmd + "\n"); err != nil {
		fmt.Println("Can't update History", err.Error())
		return
	}
}

// Interface with common methods for each command.
type Command interface {
	// Checks if the command is valid.
	IsValid() bool
	// Executes the command.
	Execute() (string, error)
}

// Command specific structs
type Ls struct{ words []string }
type Pwd struct{ words []string }
type Cd struct{ words []string }
type Exit struct{ words []string }
type History struct{ words []string }
type Help struct{ words []string }

func (l Ls) IsValid() bool {
	return len(l.words) == 1 && l.words[0] == ls
}

func (l Ls) Execute() (string, error) {
	out, err := exec.Command(l.words[0]).Output()
	return string(out), err
}

func (p Pwd) IsValid() bool {
	return len(p.words) == 1 && p.words[0] == pwd
}

func (p Pwd) Execute() (string, error) {
	out, err := exec.Command(p.words[0]).Output()
	return string(out), err
}

func (c Cd) IsValid() bool {
	return len(c.words) == 2 && c.words[0] == cd
}
func (c Cd) Execute() (string, error) {
	err := os.Chdir(c.words[1])
	return "", err
}

func (e Exit) IsValid() bool {
	return len(e.words) == 1 && e.words[0] == exit
}

func (e Exit) Execute() (string, error) {
	return exit, nil
}

func (h Help) IsValid() bool {
	return len(h.words) == 1 && h.words[0] == help
}
func (h Help) Execute() (string, error) {
	helpstr := `
	ls : list of dir/files in the current dir
	pwd : Gives current working dir
	exit : Exit the shell
	cd <destination> : Go to destination dir
	history : Display History of the commands used
	help : describes all the available commands in the shell
	`
	return helpstr, nil
}

func (h History) IsValid() bool {
	return len(h.words) == 1 && h.words[0] == history
}

func (h History) Execute() (string, error) {
	return readHistory()
}

// Function that takes Command Interface and calls the specific IsValid method.
func isValid(cmd Command) bool {
	return cmd.IsValid()
}

// Function that takes Command Interface and calls the specific Execute method.
func execute(cmd Command) (string, error) {
	return cmd.Execute()
}

// Factory funtion that returns Command Struct based on Command type
func CommandFactory(words []string) Command {
	if words[0] == ls {
		return Ls{words: words}
	} else if words[0] == exit {
		return Exit{words: words}
	} else if words[0] == pwd {
		return Pwd{words: words}
	} else if words[0] == cd {
		return Cd{words: words}
	} else if words[0] == help {
		return Help{words: words}
	} else if words[0] == history {
		return History{words: words}
	}
	return nil
}

func main() {
	fmt.Println("Go Shell started")
	scanner := bufio.NewScanner(os.Stdin)

	for i := ""; i != exit; {
		user, err := user.Current()
		if err != nil {
			fmt.Println(err.Error())
		}
		host, err := os.Hostname()
		if err != nil {
			fmt.Println(err.Error())
		}
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		wd = strings.Replace(wd, user.HomeDir, "~", 1)
		fmt.Printf("%s@%s:%s$ ", user.Username, host, wd)

		scanner.Scan()
		err = scanner.Err()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		in := scanner.Text()
		updateHistory(in)
		words := strings.Fields(in)
		var cmd Command
		if len(words) > 0 {
			cmd = CommandFactory(words)
			if cmd != nil && isValid(cmd) {
				i, err = execute(cmd)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(i)
			} else {
				fmt.Println("Invalid Command, type 'help' to learn more")
			}
		} else {
			fmt.Println("Invalid Command, type 'help' to learn more")
		}
	}
}
