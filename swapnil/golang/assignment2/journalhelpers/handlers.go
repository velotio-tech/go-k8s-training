package journalhelpers

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func loginHandler(ud userData, args []string) bool {
	if len(args) != 3 {
		fmt.Println("invalid params for login")
		return false
	} else if ud.userExists(args[1]) == false {
		fmt.Printf("Error: user %v does not exist\n", args[1])
		return false
	} else if args[2] != ud[args[1]].Password {
		fmt.Println("Error: invalid credentials")
		return false
	} else {
		fmt.Println("Info: logged in")
		return true
	}
}

func signUpHandler(ud userData, args []string) bool {
	if len(args) != 3 {
		fmt.Println("invalid params for signing up")
		return false
	} else if ud.userExists(args[1]) == true {
		fmt.Printf("Error: user %v already exist\n", args[1])
		return false
	} else if len(args[2]) < 5 {
		fmt.Println("Error: password should be atleast 5 char long")
		return false
	} else {
		result := ud.addUser(args[1], args[2])
		if result == true {
			fmt.Println("Info: logged in")
			return true
		}
		return false
	}
}

func loginSuccessMsg() {
	fmt.Printf("For adding new entry to your journal enter 'add <text to add>'\nTo see all journals enter 'list'\nFor logging out enter 'logout'\nFor quitting enter 'quit'\nTo see this help enter'help'\n")
}

func appHandler(ud userData, username string) bool {
	loginSuccessMsg()
	for {
		args := getInput()
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "add":
			text := strings.Join(args[1:], " ")
			if len(text) > 0 {
				ud.addEntry(username, entry{Text: text, CreatedAt: time.Now()})
			} else {
				fmt.Println("Error: Nothing to add")
			}

		case "list":
			ud.listEntries(username)
		case "logout":
			fmt.Println("Info: logged out")
			return true
		case "quit":
			ud.storeUserData()
			os.Exit(0)
		case "help":
			loginSuccessMsg()
		default:
			fmt.Printf("Error: that meant nothing to me\n check usage with 'help'\n")
		}
	}
}
