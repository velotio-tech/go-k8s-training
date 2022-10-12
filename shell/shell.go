package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/golang-collections/go-datastructures/queue"
)

type shell struct {
	username     string
	hostname     string
	pwd          string
	scanner      *bufio.Scanner
	historyQueue *queue.Queue
}

func New() (*shell, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	return &shell{
		username:     currentUser.Username,
		hostname:     hostname,
		pwd:          pwd,
		scanner:      scanner,
		historyQueue: queue.New(10),
	}, nil
}

func (s *shell) Start() error {
	fmt.Printf("\n%s@%s %s $ ", s.username, s.hostname, s.pwd)
	s.scanner.Scan()
	command := s.scanner.Text()
	commands := strings.Split(command, " ")
	if len(commands) == 0 {
		return nil
	}
	defer s.historyQueue.Put(command)
	return s.Run(commands[0], commands[1:]...)
}

func (s *shell) Run(command string, args ...string) error {
	switch command {
	case "ls":
		result, err := ls(args...)
		if err != nil {
			return fmt.Errorf("%s: %s", command, err.Error())
		}
		fmt.Print(result)
		return nil
	case "pwd":
		result, err := pwd(args...)
		if err != nil {
			return fmt.Errorf("%s: %s", command, err.Error())
		}
		fmt.Print(result)
		return nil
	case "exit":
		_, err := exit(args...)
		if err != nil {
			return fmt.Errorf("%s: %s", command, err.Error())
		}
		return nil
	case "cd":
		result, err := cd(args...)
		if err != nil {
			return fmt.Errorf("%s: %s", command, err.Error())
		}
		s.pwd = result
		return nil
	case "history":
		items, err := s.historyQueue.Get(10)
		if err != nil {
			return err
		}
		for _, item := range items {
			fmt.Println(item)
		}
	default:
		return fmt.Errorf("command not found : %s", command)
	}
	return nil
}
