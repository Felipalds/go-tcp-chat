package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// pad applies PKCS#7 padding to a plaintext.
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(src, padtext...)
}

// unpad removes PKCS#7 padding from a plaintext.
func unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// encryptAES256ECB encrypts plaintext using AES-256 in ECB mode.
func Encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(plaintext)%aes.BlockSize != 0 {
		plaintext = pad(plaintext, aes.BlockSize)
	}

	ciphertext := make([]byte, len(plaintext))
	for bs, be := 0, block.BlockSize(); bs < len(plaintext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(ciphertext[bs:be], plaintext[bs:be])
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// decryptAES256ECB decrypts ciphertext using AES-256 in ECB mode.
func Decrypt(ciphertextBase64 string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for bs, be := 0, block.BlockSize(); bs < len(ciphertext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(plaintext[bs:be], ciphertext[bs:be])
	}

	return string(unpad(plaintext)), nil
}

// Encrypt encrypts plain text string into cipher text string
func EncryptOld(plainText string, keyBytes []byte) (string, error) {
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
func DecryptOld(cipherText string, keyBytes []byte) (string, error) {
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
func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32) // 32 bytes for AES-256
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err.Error())
	}

	return key, nil
}
