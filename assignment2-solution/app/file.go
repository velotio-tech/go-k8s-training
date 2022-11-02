package app

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var file string = "data.txt"
var uname string
var splitName []string
var flag bool

func AddNewEntryToFile(username string) {
	fmt.Println("Saving data to file")
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error : open operation failed!")
	}

	var uname, newEntryIntoFile string
	check, err := f.Stat()
	if check.Size() != 0 {
		fmt.Println("Non Empty file")

		ff, err := os.ReadFile(file)
		fileData := string(ff)
		if err != nil {
			fmt.Println("Error : open operation failed!")
		}

		fmt.Println(fileData)

		if len(fileData) == 0 {
			return
		}
		arr := strings.Split(fileData, "###")
		for i, val := range arr {
			if len(val) <= 1 {
				continue
			}

			splitName := strings.Split(arr[i], "/")

			uname = splitName[0]
			splitUname := strings.Split(uname, "---")
			uname = splitUname[0]

			if uname == username {
				fmt.Println("Breaking here?")
				flag = true
				break
			}
		}

	}

	for key, value := range userMap {

		var str = ""
		var i int
		for i = 0; i < len(value.time); i++ {
			str = str + value.time[i] + " / " + value.data[i] + " / "
		}

		if uname != "" && flag == true {
			splitKey := strings.Split(uname, "---")
			key = splitKey[0]
		}

		if uname == key && flag == true {
			fmt.Println("Entry exists in a file")
			newEntryIntoFile = "---" + str
			//break
		} else {
			newEntryIntoFile = "\n###" + key //+ "---" + str
		}

	}
	//EncryptData()
	_, err = f.Write([]byte(newEntryIntoFile))
	if err != nil {
		log.Fatal(err)
		f.Close()
		return
	}

	fmt.Println("Entry successfully added!")
}

func GetDataFromFile() {
	// decrypt the data
	f, err := os.ReadFile("data.txt")
	fileData := string(f)
	if err != nil {
		fmt.Println("Error : open operation failed!")
	}

	fmt.Println(fileData)

	if len(fileData) == 0 {
		return
	}
	arr := strings.Split(fileData, "###")
	for i, val := range arr {
		if len(val) <= 1 {
			continue
		}

		splitName = strings.Split(arr[i], "/")

		uname := splitName[0]
		var data EntryData
		userMap[uname] = data
		List()
	}

}

func List() {
	var i int
	fmt.Println("Printing usermap here :")
	for i = 1; i < len(splitName)-1; i += 2 {

		//AppendValueToMap(key, arr2[i], arr2[i+1])
		fmt.Println("Entry : ", splitName[i])
		user, ok := userMap[uname]
		if ok {
			user.time = append(user.time, splitName[i])
			user.data = append(user.data, splitName[i+1])

		}

		userMap[uname] = user
	}
}

var secretData = []byte("abc&1*~#^2^#s0^=)^^7%b34")

func EncryptFile() {
	fdata, _ := ioutil.ReadFile(file)
	data := string(fdata)
	e := os.Remove("data.txt")
	if e != nil {
		fmt.Println("Error: Couldn't able to remove the file")
	}
	filename, _ := os.Create(file)

	encryptdata := Encrypt(data)
	_, err := filename.Write([]byte(encryptdata))
	if err != nil {
		fmt.Println("Error: Not able to write data")
		filename.Close()
		return
	}
}

func Encrypt(message string) string {
	plainText := []byte(message)

	// AES cipher using secret
	block, err := aes.NewCipher(secretData)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText)
}

func DecryptFile() {
	fdata, _ := ioutil.ReadFile(file)
	data := string(fdata)
	e := os.Remove(file)
	if e != nil {
		fmt.Println("Error occured")
	}
	filename, _ := os.Create(file)

	encryptdata := Decrypt(data)
	_, err := filename.Write([]byte(encryptdata))
	if err != nil {
		fmt.Println("Error occured during write")
		filename.Close()
		return
	}
}

func Decrypt(data string) string {
	// Remove base64 encoding
	cipherText, err := base64.RawStdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(secretData)
	if err != nil {
		panic(err)
	}

	// Return error, if length of the cipherText is less than 16 Bytes
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return ""
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the file
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
