package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	internalCrypto "github.com/AmitKarnam/KeyCloak/internal/utils/crypto"
)

// AES Type which implements methods of crypto type
type AES struct{}

var _ internalCrypto.Crypto = &AES{}

// Encrypt method for AES
func (a *AES) Encrypt(key, plainText string) (string, error) {

	if key == "" {
		return "", errors.New("empty key not allowed")
	}

	if len(key) < internalCrypto.KeySize {
		return "", errors.New("key too short")
	}

	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plainText))

	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt method for AES
func (a *AES) Decrypt(key, cipherHex string) (string, error) {

	if key == "" {
		return "", errors.New("empty key not allowed")
	}

	if len(key) < internalCrypto.KeySize {
		return "", errors.New("key too short")
	}

	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("Crypto: cipher too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)
	return fmt.Sprintf("%s", cipherText), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := sha256.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
