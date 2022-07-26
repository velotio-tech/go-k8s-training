package utils

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var secret = []byte("abc&1*~#^2^#s0^=)^^7%b34")
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

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

func Encrypt_(text string) string {
	plaintext := []byte(text)
	block, _ := aes.NewCipher(secret)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error in Encryption")
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return string(ciphertext)
}

func Decrypt_(line string) string {
	ciphertext := []byte(line)
	block, err := aes.NewCipher(secret)
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

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func EntryCheck(username string) {
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
	var err error
	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	journalDir := homeDir + "/journal/"
	fileName := username + ".txt"
	userDir := journalDir + fileName
	if _, err := os.Stat(userDir); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("User: [%s] don't have journals!!\n\n\n", username)
		return
	}
	var ciphertext []byte

	ciphertext, err = ioutil.ReadFile(userDir)
	if err != nil {
		fmt.Println("Error while reading user journal, err: ", err)
	}
	lines := strings.Split(string(ciphertext), "\n")
	fmt.Println("\nListing User journal:")
	for _, v := range lines {
		if v == "" {
			break
		}
		line := Decrypt(v)
		fmt.Println(line)
	}
	fmt.Println()
}

func AddNewEntry(username string) {
	var err error
	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	journalDir := homeDir + "/journal/"
	fileName := username + ".txt"
	userDir := journalDir + fileName
	if !UserLimitCheck(userDir, 50) {
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter text for the new entry: ")
	reader := bufio.NewReader(os.Stdin)
	entry, _ := reader.ReadString('\n')
	entry = strings.TrimSuffix(entry, "\n")
	timeStamp := time.Now().Format("02-Jul-06 03:04pm")
	newEntry := Encrypt(timeStamp + "\t" + entry)
	file, err := os.OpenFile(userDir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error while opening Users file, err: ", err)
		return
	}
	defer file.Close()
	if _, err = file.WriteString(newEntry + "\n"); err != nil {
		fmt.Println("Error while adding user entry to file, err: ", err)
		return
	}
	fmt.Println("Successfully added new entry to the file!!\n")
}

func UserLimitCheck(filename string, limit int) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return true
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening users file!")
		return false
	}
	defer file.Close()
	var res int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res++
	}
	if res >= limit {
		fmt.Println("User entries limit exceeded!!")
		return false
	}
	return true
}
