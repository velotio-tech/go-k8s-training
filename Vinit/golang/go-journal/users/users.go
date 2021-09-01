package users

import (
	"fmt"
	"go-journal/constants"
	//"fmt"
	"go-journal/files"
)

type User struct {
	Username string
	password string
	IsAuthenticated bool
}

func AlreadyExists(user, pass, chk string) bool {
	userMap := files.GetUserList()
	if val, ok := userMap[user]; ok {
		if chk == "signup" {
			return true
		}
		if chk == "login" {
			if pass==val {
				return true
			} else {
				return false
			}
		}
		return false
	} else {
		return false
	}
}

func CreateNew (username, pass string) *User {
	var totalUsers, _ = files.GetTotalUsers()

	if totalUsers <= constants.MAX_USERS {
		files.AddToUsers(username, pass)
		files.CreateJournal(username + constants.JOURNAL_NAME)
		newUser := User{
			Username:        username,
			password:        pass,
			IsAuthenticated: true,
		}
		return &newUser
	} else {
		fmt.Println("User Limit exceeded, cannot create a new user")
		return nil
	}
}

func (u *User) ReadJournal() {
	files.ReadJournalFile(constants.DB_LOCATION+u.Username+constants.JOURNAL_NAME)
}

func (u *User) WriteJournal (journal string){
	files.WriteJournalFile(constants.DB_LOCATION+u.Username+constants.JOURNAL_NAME, journal)
}

func GetData(username string) *User {
	var loggedInUser = User{
		Username: username,
		password: "",
		IsAuthenticated: true,
	}
	return &loggedInUser
}