package files

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"go-journal/constants"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func encrypt(data string) string {
	plaintext := []byte(data)
	key := []byte(constants.ENCRYPTION_KEY)

	block, err := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error in Encryption")
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func decrypt(cipherstring string) string {
	ciphertext := []byte(cipherstring)
	key := []byte(constants.ENCRYPTION_KEY)
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

func GetUserList() map[string]string {
	var userToPass = map[string]string{}
	_, cipherUser := GetTotalUsers()
	if cipherUser == nil || len(cipherUser) < 1 {
		return userToPass
	}
	lines := strings.Split(string(cipherUser), "\n")
	for _, v := range lines {
		if v == ""{
			break
		}
		plainLine := decrypt(v)
		uPmap := strings.Split(plainLine, ":")
		userToPass[uPmap[0]] = uPmap[1]
	}
	return userToPass
}

func AddToUsers(username, pass string) {
	file := constants.DB_LOCATION+constants.USER_DB
	entry := username+":"+pass
	writeFile(entry, file)
}

func writeFile(data, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		_, err = f.Write([]byte(encrypt(data) + "\n"))
	//_, err = f.Write([]byte(data + "\n"))
	if err!= nil {
		fmt.Println("Error in adding to file")
		return
	}
	f.Close()
}

func readFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func ReadJournalFile(file string) []string {
	var plaintext []string
	if ciphertext, err := readFile(file); err != nil {
		fmt.Println("File is not found")
		return make([]string, 0)
	} else {
		if len(ciphertext) < 1 {
			return plaintext
		}
		fmt.Println("^^^^^^^^^^^^^^^^^^", string(ciphertext))
		lines := strings.Split(string(ciphertext), "\n")
		for _, v := range lines {
			if v == ""{
				break
			}
			plaintext = append(plaintext, decrypt(v))
		}
		fmt.Println(plaintext)
		return plaintext
	}
}

func CreateJournal(journal string){
	userJournal, e := os.Create(constants.DB_LOCATION+journal)
	if e != nil {
		fmt.Println("Error Creating a blank journal")
	}
	userJournal.Close()
}

func GetTotalUsers() (int, []byte) {
	if ciphertext, err := readFile(constants.DB_LOCATION+constants.USER_DB); err != nil {
		fmt.Println("File is not found", err)
		return 11, nil
	} else {
		lines := strings.Split(string(ciphertext), "\n")
		return len(lines), ciphertext
	}
}

func WriteJournalFile(file string, journal string) {
	data := ReadJournalFile(file)
	var entry string
	var writeJournal = make([]string, 50, 50)
	writeJournal = append(writeJournal, data...)
	writeJournal = append(writeJournal, journal)
	if len(writeJournal) > constants.MAX_JOURNAL_ENTERIES{
		diff := len(writeJournal) - constants.MAX_JOURNAL_ENTERIES
		writeJournal = writeJournal[diff:]
	}
	for _, v := range writeJournal{
		entry += v+"\n"
	}
	writeFile(entry, file)
}