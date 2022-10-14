package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/spf13/viper"
)

func Encrypt(src any) (out string, err error) {
	plaintext, err := json.Marshal(src)
	if err != nil {
		return
	}
	key := viper.GetString("INIT_SECRET")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	cipherText := make([]byte, aes.BlockSize+len(plaintext))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	//Return string encoded in base64
	return base64.RawStdEncoding.EncodeToString(cipherText), err
}
