package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var FILE_NAME string

type Entry struct {
	Time []string
	Info []string
}

var UserEntryLimit = 5

// var my_map = map[string][]entry{}
var My_map = make(map[string]Entry)

// var set = make(map[string]void)
var secret = []byte("abc&1*~#^2^#s0^=)^^7%b34")

func Encrypt(message string) string {
	plainText := []byte(message)

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	// Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText)
}

func Decrypt(secure string) string {
	// Remove base64 encoding
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	// If the length of the cipherText is less than 16 Bytes
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return ""
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

func EntryCheckLoop(username string) {
	for {
		fmt.Print("1.Add new entry\n2.List all entries\n3.Exit\nEnter Your Option: ")
		var entry string
		fmt.Scanln(&entry)
		switch entry {
		case "1":
			AddNewEntry(username)
		case "2":
			ListAllEntries(username)
		case "3":
			return
		default:
			fmt.Println("Please enter correct option.")
		}
	}
}

func ListAllEntries(username string) {
	user, ok := My_map[username]
	if ok {
		var i int
		for i = 0; i < len(user.Time); i++ {
			fmt.Println(user.Time[i] + "\t" + user.Info[i])
		}
	}

}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func AddNewEntry(username string) {

	fmt.Println("Enter the data to store")
	var data string
	fmt.Scanln(&data)
	//var obj Entry
	currentTime := time.Now()
	timeInShort := currentTime.Format("01-02-2006 15:04:05 Monday")
	//obj.AppendValueToMap(currentTime.String(), data)
	//	My_map[username] = obj

	AppendValueToMap(username, timeInShort, data)
	fmt.Println("Successfully added new entry to the file!!\n")
}

func AppendValueToMap(user, time, data string) {

	username, ok := My_map[user]
	if ok {
		if len(username.Time) == UserEntryLimit {
			username.Time = RemoveIndex(username.Time, 0)
			username.Info = RemoveIndex(username.Info, 0)
		}
		username.Time = append(username.Time, time)
		username.Info = append(username.Info, data)

		//fmt.Println("Value insertion sucessful")
	} else {
		fmt.Println("Value insertion failed")
	}
	My_map[user] = username
}

func LoginOrSignUpLoop() {
	for {
		fmt.Print("1.for login \n2.sign up\n3.Exit\nEnter Your Option: ")
		var entry string
		fmt.Scanln(&entry)
		switch entry {
		case "1":
			LoginUser()
		case "2":
			SignUpUser()
		case "3":
			SaveDataToFileHelper()
			os.Exit(0)
		default:
			fmt.Println("Please enter correct option.")
		}
	}
}

func LoginUser() {

	fmt.Println("Enter the username")
	var username string
	fmt.Scanln(&username)
	LoginHelper(username)

	EntryCheckLoop(username)
}

func SignUpUser() {

	if len(My_map) == 10 {
		fmt.Println("User Account limit exceeded")
		return
	}

	fmt.Println("Enter the username")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Enter the password")
	var password string
	fmt.Scanln(&password)
	SignupHelper(username, password)
	EntryCheckLoop(username)

}

func LoginHelper(username string) {
	if _, ok := My_map[username]; ok {

		//SaveDataToFileHelper(username)
	} else {
		fmt.Println("Sign up require!")
	}
}

func SignupHelper(username string, password string) {
	var obj Entry
	My_map[username] = obj
	//fmt.Println(utils.My_map)
	//SaveDataToFileHelper(username)
	fmt.Println("Sign up sucessfully with id " + username + " and password is " + password)
}
