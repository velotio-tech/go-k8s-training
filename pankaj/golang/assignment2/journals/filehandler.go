package journals

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(fileName string, passphrase string) string {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return ""
	}

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	data = decrypt(data, passphrase)
	return string(data)

}

func WriteFile(filepath string, data string, passphrase string) {
	file, _ := os.Create(filepath)
	defer file.Close()
	file.Write(encrypt([]byte(data), passphrase))
}
