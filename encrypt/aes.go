package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypt encrypts plain text string into cipher text string
func Encrypt(plainText string, keyBytes []byte) (string, error) {
	plaintextBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, plaintextBytes, nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(cipherText string, keyBytes []byte) (string, error) {
	fmt.Println("Cipher! ", cipherText)
	fmt.Println("Key Bytes! ", keyBytes)
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherTextBytes) < nonceSize {
		return "", fmt.Errorf("cipherText too short")
	}
	nonce, cipherTextBytes := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]

	plainTextBytes, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainTextBytes), nil
}
