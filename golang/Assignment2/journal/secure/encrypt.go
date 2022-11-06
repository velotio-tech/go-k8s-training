package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/spf13/viper"
)

// Ref : https://www.golinuxcloud.com/golang-encrypt-decrypt/

func Encrypt(bytesToEncrypt []byte) (string, error) {
	key := viper.GetString("key")
	//Create a new Cipher Block from the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(bytesToEncrypt))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], bytesToEncrypt)

	return base64.RawStdEncoding.EncodeToString(ciphertext), nil
}
