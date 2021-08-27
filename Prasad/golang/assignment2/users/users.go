package users

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/constants"
	"github.com/thisisprasad/go-k8s-training/Prasad/golang/assignment2/encdec"
)

//	private struct. Managed by users package.
type loggedInUser struct {
	username string
}

var activeUser *loggedInUser

type user struct {
	username string
	passwd   string
}

var users []user

//	Returns the logged-in user.
func LoggedInUser() (string, error) {
	if activeUser == nil {
		return "", errors.New("no logged-in user")
	}
	return activeUser.username, nil
}

//	logs-in a user.
func Login(uname string) {
	activeUser = &loggedInUser{username: uname}
}

//	logs-out a user if logged-in.
func Logout() {
	activeUser = nil
}

//	Look for user in registry and cross-check the credentials
func ValidateUser(uname string, passw string) bool {
	for _, user := range users {
		if user.username == uname && user.passwd == passw {
			return true
		}
	}
	return false
}

//	Initializes the users data
func InitUserData() error {
	err := createUsersFileIfNotExists()
	loadUsersData()
	return err
}

func createUsersFileIfNotExists() error {
	//	Check if folder exists or not.
	userfp := constants.UserRegistryPath
	_, err := os.Stat(userfp)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(userfp, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else if err != nil {
		return err
	}
	//	Check if file exists or not in the path.
	userfp = constants.UserRegistryPath + constants.UserRegistryFile
	_, err = os.Stat(userfp)
	if errors.Is(err, os.ErrNotExist) {
		_, ferr := os.Create(userfp)
		if ferr != nil {
			return ferr
		}
	} else if err != nil {
		return err
	}

	return nil
}

func loadUsersData() error {
	userfp := constants.UserRegistryPath + constants.UserRegistryFile
	file, err := os.OpenFile(userfp, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	rawData, err := ioutil.ReadFile(userfp)
	if err != nil {
		return err
	} else if len(rawData) == 0 {
		return nil
	}
	decBuffer := encdec.Decrypt(rawData)
	strData := string(decBuffer)
	usersData := strings.Split(strData, constants.JournEntryDelim)
	for _, record := range usersData {
		urecord := strings.Split(record, constants.JournValDelim)
		user := user{
			username: urecord[0],
			passwd:   urecord[1],
		}
		users = append(users, user)
	}

	return nil
}

//	Creates a new user.
//	Adds an entry in users file
func CreateNewUser(uname string, passw string) bool {
	if getuserCount() >= constants.MaxUserLimit {
		fmt.Println("Error: max user limit reached.")
		return false
	}

	newUser := user{
		username: uname,
		passwd:   passw,
	}
	users = append(users, newUser)
	// var sb strings.Builder
	// for i := 0; i < len(users); i++ {
	// 	sb.WriteString(users[i].username + constants.JournValDelim + users[i].passwd)
	// 	if i != len(users)-1 {
	// 		sb.WriteString(constants.JournEntryDelim)
	// 	}
	// }
	// plainText := sb.String()
	// cipherText := encdec.Encrypt(plainText)

	writeUsersDataToFile()

	// userfp := constants.UserRegistryPath + constants.UserRegistryFile
	// file, err := os.OpenFile(userfp, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	log.Fatalln("Error:", err)
	// 	return false
	// }
	// defer file.Close()

	// fwriter := bufio.NewWriter(file)
	// if _, err := fwriter.Write(cipherText[:]); err != nil {
	// 	fmt.Println("Error creating user ::", err)
	// 	return false
	// }

	// fwriter.Flush()

	return true
}

func writeUsersDataToFile() error {
	var sb strings.Builder
	for i := 0; i < len(users); i++ {
		sb.WriteString(users[i].username + constants.JournValDelim + users[i].passwd)
		if i != len(users)-1 {
			sb.WriteString(constants.JournEntryDelim)
		}
	}
	cipherText := encdec.Encrypt([]byte(sb.String()))

	userfp := constants.UserRegistryPath + constants.UserRegistryFile
	file, err := os.OpenFile(userfp, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	fwriter := bufio.NewWriter(file)
	if _, err := fwriter.Write(cipherText[:]); err != nil {
		return err
	}
	fwriter.Flush()

	return nil
}

//	checks whether user with newUname already exists
func UserAlreadyExists(newUname string) bool {
	for _, user := range users {
		if user.username == newUname {
			return true
		}
	}
	return false
}

func getuserCount() int {
	return len(users)
}
