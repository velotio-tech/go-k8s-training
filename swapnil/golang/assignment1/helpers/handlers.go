package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

func HandleLS(args []string) {
	// dirname wil be showed if multiple paths are given to ls
	showDirName := false
	if len(args) == 1 {
		// directory is not given, taking current working directory as default
		cwd := GetCWD()
		args = append(args, cwd)
	} else if len(args) > 2 {
		showDirName = true
	}

	for _, path := range args[1:] {
		path = replaceSymbolWithHomeFolder(path)
		if showDirName == true {
			fmt.Println(path + ":\n")
		}

		fs, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("Error: ", err)
			// if error occures for one path try next path
			continue
		}
		for _, each := range fs {
			resetColor := ""
			blueColor := ""
			if each.IsDir() == true {
				resetColor = "\033[0m"
				blueColor = "\033[34m"
			}
			// directory will be showed in blue color
			fmt.Println(blueColor + each.Name() + resetColor)

		}
	}
}

func HandlCD(args []string) {
	isValid := true
	if len(args) > 2 {
		fmt.Println("Error: too many arguments")
		isValid = false
	} else if len(args) == 1 {
		// user didn't provide "path" to cd into, default is home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		args = append(args, homeDir)
	}
	if isValid == true {
		// if user uses ~ in path replace with actual home path
		path := replaceSymbolWithHomeFolder(args[1])
		err := os.Chdir(path)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
