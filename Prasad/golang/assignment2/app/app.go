package app

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/journal"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/users"
)

//	Starts the cmd-client of journal application for uname user.
func LaunchCMDclient(uname string) {
	fmt.Print("\033[H\033[2J")
	startCMDapp(uname)
}

func startCMDapp(uname string) {
	var choice int
	var err error
	scheduleInterruptRoutine()
	for {
		fmt.Println("1. Load journal entries")
		fmt.Println("2. Add a new entry.")
		fmt.Println("3. Exit")
		fmt.Print("Choice: ")
		fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Error:", err)
		}

		switch choice {
		case 1:
			showJournData()
			break

		case 2:
			addJournalEntry()
			break

		case 3:
			CloseJournApp()
			os.Exit(0)

		default:
			fmt.Println("Invalid choice...")
			break
		}
	}
}

//	This functions starts a routine to capture SIGTERM interrrupt(Ctrl+C).
//	to perform cleanup operations and to ensure that program exits gracefully.
func scheduleInterruptRoutine() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		CloseJournApp()
		os.Exit(1)
	}()
}

func showJournData() {
	journData := journal.GetUserJournData()
	logdInUser, _ := users.LoggedInUser()
	fmt.Println("User:", logdInUser)
	for _, entry := range journData {
		fmt.Println("==>", entry)
	}
	fmt.Println()
}

//	adds an entry into the journal.
func addJournalEntry() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Entry -> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	je := &journal.JournalEntry{
		Timestamp: time.Now().Format("02 Jan 2006 15:04"),
		Text:      text,
	}
	journal.AddEntry(*je)
}

//	logs-in a user.
func LoginUser(uname string) {
	users.Login(uname)
	journal.LoadUserJournalData(uname)
}

//	Closes the CMD client of the application.
func CloseJournApp() {
	journal.WriteJournDataToFile()
	users.Logout()
}
