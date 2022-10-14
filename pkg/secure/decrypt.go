package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/spf13/viper"
)

func Decrypt(src string, dst any) (err error) {
	cipherText, err := base64.RawStdEncoding.DecodeString(src)
	if err != nil {
		return
	}
	key := viper.GetString("INIT_SECRET")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	err = json.Unmarshal(cipherText, &dst)
	return
}
