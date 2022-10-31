package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const (
	ls   string = "ls"
	pwd         = "pwd"
	cd          = "cd"
	exit        = "exit"
	help        = "help"
)

// Tried to use factory pattern but didn't implement it, which would require me to
type Command interface {
	IsValid() bool
	Execute() (string, error)
}

type Ls struct{ words []string }
type Pwd struct{ words []string }
type Cd struct{ words []string }
type Exit struct{ words []string }
type Help struct{ words []string }

func (l Ls) IsValid() bool {
	return len(l.words) == 1 && l.words[0] == ls
}

// 	out, err := exec.Command("ls").Output()
// 	out2, err := exec.Command("pwd").Output()
// 	out3, err := exec.Command("cd", "../").Output()
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
	out, err := exec.Command("bash", "-c", fmt.Sprintf("%s %s", c.words[0], c.words[1])).Output()
	return string(out), err
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
	exit : Exit the CLI
	cd <destination> : Go to destination dir
	`
	return helpstr, nil
}

func isValid(cmd Command) bool {
	return cmd.IsValid()
}

func execute(cmd Command) (string, error) {
	return cmd.Execute()
}

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
	}
	return nil
}

func main() {
	fmt.Println("Go CLI started")
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
		words := strings.Fields(in)
		var cmd Command
		if len(words) > 0 {
			cmd = CommandFactory(words)
			if isValid(cmd) {
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
