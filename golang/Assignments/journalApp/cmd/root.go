package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var HOME_PATH string = os.Getenv("HOME") + "/.journalApp"

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "CLI for adding journal entries and user management",
	Long:  `CLI for adding journal entries and user management`,
	PreRun: func(cmd *cobra.Command, args []string) {
		path := os.Getenv("HOME") + "/.journalApp"
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0700)
			if err != nil {
				log.Fatalf("Mkdir %q: %s", path, err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Journal App")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Press 1 for login.\nPress 2 for signup.")
		scanner.Scan()
		choice := scanner.Text()
		if choice == "1" {
			username, password := getUserCreds(scanner)
			if login(username, password) {
				fmt.Println("Log in successfull")
				showDashboard(scanner, username, password)
			} else {
				fmt.Println("Invalid Credentials. Exiting Program..")
				os.Exit(1)
			}
		} else if choice == "2" {
			username, password := getUserCreds(scanner)
			signup(username, password)
			showDashboard(scanner, username, password)
		}

	},
}

func getUserCreds(scanner *bufio.Scanner) (string, string) {
	fmt.Println("Enter username")
	scanner.Scan()
	username := scanner.Text()
	fmt.Println("Enter password")
	scanner.Scan()
	password := scanner.Text()
	return username, password
}
func showDashboard(scanner *bufio.Scanner, username string, password string) {

	for {
		fmt.Println("Press 1 for seeing old journal entry.\nPress 2 for adding new journal entry.")
		fmt.Println("Press 3 to exit the application")
		scanner.Scan()
		if scanner.Text() == "1" {
			showJournal(username, password)
		} else if scanner.Text() == "2" {
			fmt.Println("Enter the new text for journal")
			scanner.Scan()
			addJournalEntry(username, password, scanner.Text())
		} else if scanner.Text() == "3" {
			os.Exit(1)
		} else {
			fmt.Println("Invalid input")
		}
	}
}

func showJournal(username string, password string) {
	fmt.Println(string(decryptFile(HOME_PATH+"/"+username+"/"+"journal", password)))
}
func addJournalEntry(username string, password string, data string) bool {
	journalText := string(decryptFile(HOME_PATH+"/"+username+"/"+"journal", password))
	textData := journalText + string(time.Now().Format(time.RFC850)) + "- " + data + "\n"
	encryptFile(HOME_PATH+"/"+username+"/"+"journal", []byte(textData), password)
	return true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
