/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package journalhelpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() []string {
	fmt.Printf("$ ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// removing last new line char
	text = text[:len(text)-1]
	// get slice of words by splitting on whitespace
	args := strings.Fields(text)
	if len(args) == 0 {
		//nothing is typed so continuing
		fmt.Println("try again")
	}
	return args
}

// Main is myjournal app entrypoint
func Main() {
	data := loadUserData()
	fmt.Println(getHome())
	for {
		args := getInput()
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "quit":
			data.storeUserData()
			os.Exit(0)
		case "login":
			result := loginHandler(data, args)
			if result == false {
				continue
			} else {
				result := appHandler(data, args[1])
				if result == true {
					continue
				}
			}
		case "signup":
			result := signUpHandler(data, args)
			if result == false {
				continue
			} else {
				result := appHandler(data, args[1])
				if result == true {
					continue
				}
			}
		case "help":
			fmt.Println(getHome())
		default:
			fmt.Printf("Error: that meant nothing to me\n check usage with 'help'\n")
		}
	}
}
