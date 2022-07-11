package service

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func GetChoice(username string) {
	for {
		fmt.Println("1.Add new entry\n2.List all entries\n3.Exit\nEnter Your Option: ")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			AddNewEntry(username)
		case "2":
			ListAllEntries(username)
		case "3":
			return
		default:
			fmt.Println("Please enter correct option")
		}
	}
}

func ListAllEntries(username string) {
	// Check if user entries file exist or not
	if _, err := os.Stat("./entries/" + username + ".txt"); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("User: [%s] don't have journals!!\n", username)
		return
	}
	var ciphertext []byte
	var err error
	if ciphertext, err = ioutil.ReadFile("./entries/" + username + ".txt"); err != nil {
		fmt.Println("Error while reading user journal, err: ", err)
	}
	lines := strings.Split(string(ciphertext), "\n")
	fmt.Println("User journal:")
	// Display all entries from the user journal
	for _, v := range lines {
		if v == "" {
			break
		}
		line := Decrypt(v)
		fmt.Println(line)
	}
}

func AddNewEntry(username string) {
	// Check if user entries limit exceeded
	if !checkForLimit("./entries/"+username+".txt", 50) {
		return
	}
	fmt.Println("Enter some text for new entry: ")
	// Scan new user entry for journal
	reader := bufio.NewReader(os.Stdin)
	entry, _ := reader.ReadString('\n')
	entry = strings.TrimSuffix(entry, "\n")
	timeStamp := time.Now().Format("06-Jan-02 03:04pm")
	newEntry := Encrypt(timeStamp + "\t" + entry)
	file, err := os.OpenFile("./entries/"+username+".txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error while opening Users file, err: ", err)
		return
	}
	defer file.Close()
	if _, err = file.WriteString(newEntry + "\n"); err != nil {
		fmt.Println("Error while adding user entry to file, err: ", err)
		return
	}
	fmt.Println("Successfully added new entry to the file!!")
}

func Encrypt(text string) string {
	plaintext := []byte(text)
	key := []byte("LetBeBestInWorld")
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error in Encryption")
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return string(ciphertext)
}

func Decrypt(line string) string {
	ciphertext := []byte(line)
	key := []byte("LetBeBestInWorld")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		fmt.Println("Error in decryption")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}
