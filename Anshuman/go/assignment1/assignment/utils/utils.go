package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

var ErrorSt string = "this version of %s does not support the arguments provided\n"

func getHistoryFileObj(forWrite bool) *os.File {
	path := "/tmp/.vhistory"
	if !forWrite {
		fileObj, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		return fileObj
	}
	if _, err := os.Stat(path); err == nil {
		fileObj, fErr := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if fErr != nil {
			panic(fErr)
		}
		return fileObj
	} else if errors.Is(err, os.ErrNotExist) {
		fileObj, fErr := os.Create(path)
		if fErr != nil {
			panic(fErr)
		}
		return fileObj
	} else {
		panic(err)
	}
}

func GetSignature() string {
	curr_user, _ := user.Current()
	working_dir := pwdHelper()
	username := curr_user.Username
	hostname, _ := os.Hostname()

	return fmt.Sprintf("%s@%s:~%s$ ", username, hostname, working_dir)
}

func PWD() {
	fmt.Println(pwdHelper())
}

func pwdHelper() string {
	working_dir, _ := os.Getwd()
	return working_dir
}

func CD(path string) {
	os.Chdir(path)
}

func History() {
	fileObj := getHistoryFileObj(false)
	defer fileObj.Close()
	scanner := bufio.NewScanner(fileObj)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func LS(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("invalid path or arguments supplied")
		return
	}
	defer file.Close()
	filesInDir, err := file.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, fileName := range filesInDir {
		fmt.Printf("\t%s\n", fileName.Name())
	}
}

func RecordHistory(cmd string) {
	fileObj := getHistoryFileObj(true)
	defer fileObj.Close()
	lineNo := getLineNumber()
	lineNo = lineNo + 1
	fileObj.WriteString(fmt.Sprintf("%d %s", lineNo, cmd))
}

func getLineNumber() int {
	fileObj := getHistoryFileObj(false)
	scanner := bufio.NewScanner(fileObj)
	text := ""
	for {
		if !scanner.Scan() {
			break
		}
		text = scanner.Text()
	}
	if text == "" {
		return 0
	}
	splitString := strings.Split(text, " ")
	num, err := strconv.Atoi(splitString[0])
	if err != nil {
		panic(err)
	}
	return num
}

func ShowValidCommands() {
	fmt.Println("the supported commands are: \nls\tpwd\texit\tcd\thistory")
}
