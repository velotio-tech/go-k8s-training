package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Session struct {
	hostname string
	os       string
	cwd      string
}

var session Session

func setup() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	session = Session{
		hostname: hostname,
		os:       runtime.GOOS,
		cwd:      cwd,
	}

	for {

		splited := strings.Split(session.cwd, "/")
		cwd = splited[len(splited)-1]

		fmt.Printf("%s@%s ~/%s ", session.hostname, runtime.GOOS, cwd)

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		args := strings.Split(strings.TrimSuffix(input, "\n"), " ")
		commandHandler(args)

	}
}
