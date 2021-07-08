package journals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func JournalInput(scanner *bufio.Scanner, userJournalPath string, journals *Journal, passphrase string) {
	defer WriteFile(userJournalPath, GetJournal(journals), passphrase)
	for {
		fmt.Printf("Journel Managment select below option.\n1. List \n2. Create New\n3. Logout\n")
		scanner.Scan()

		option, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Not a valid number")
		}
		switch option {
		case 1:
			for item := journals.Item.Back(); item != nil; item = item.Prev() {
				fmt.Printf("%v\n", item.Value)
			}
		case 2:
			fmt.Println("Enter journal details.")
			scanner.Scan()
			journal := scanner.Text()
			if journal != "" {
				journal = AddJournal(journal)
				journals.Capture(journal)
			}

		case 3:
			fmt.Println("Logout successfull!")
			WriteFile(userJournalPath, GetJournal(journals), passphrase)
			os.Exit(1)
		default:
			fmt.Println("Invalid option")
		}

	}
}

func AddJournal(entry string) string {
	curretTime := time.Now().Format("02 Jan 2006 3:4pm")
	entry = curretTime + " - " + entry
	return entry

}

func GetJournal(jouranl *Journal) string {
	var jouranls string
	for item := jouranl.Item.Front(); item != nil; item = item.Next() {
		// if len(item.Value) != 0 {
		jouranls += fmt.Sprintf("%v\n\n", item.Value)
		// }
	}
	return jouranls
}
