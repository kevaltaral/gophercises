package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Encrypt will take in a key and plaintext and return a hex representation
// of the encrypted value.
// This code is based on the standard library examples at:
//   - https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
// Encrypt will encrypt the key
func Encrypt(key, plaintext string) (string, error) {
	block, _ := newCipherBlock(key)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	_, err := io.ReadFull(rand.Reader, iv)
	if err == nil {
		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	}
	return fmt.Sprintf("%x", ciphertext), err
}

// Decrypt will take in a key and a cipherHex (hex representation of
// the ciphertext) and decrypt it.
// This code is based on the standard library examples at:
//   - https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
func Decrypt(key, cipherHex string) (string, error) {
	var ciphertext []byte
	block, err := newCipherBlock(key)
	if err == nil {
		ciphertext, err = hex.DecodeString(cipherHex)
		if err == nil {
			if len(ciphertext) >= aes.BlockSize {
				iv := ciphertext[:aes.BlockSize]
				ciphertext = ciphertext[aes.BlockSize:]

				stream := cipher.NewCFBDecrypter(block, iv)

				// XORKeyStream can work in-place if the two arguments are the same.
				stream.XORKeyStream(ciphertext, ciphertext)
			}
		}
	}
	return string(ciphertext), err
}
func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
