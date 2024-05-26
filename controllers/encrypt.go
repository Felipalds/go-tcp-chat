package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"strings"
)

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func EncodePublicToBase64(privateKey *rsa.PrivateKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", err
	}
	publicKeyPem := base64.StdEncoding.EncodeToString(publicKeyBytes)
	return publicKeyPem, nil
}

func DecryptAESKey(key string, pk *rsa.PrivateKey) {
	fmt.Println("PK", pk)
	key = strings.Replace(key, "\n", "", -1)
	fmt.Println(len(key))
	encryptedAESKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Println("Erro ao decodificar chave AES criptografada:", err)
		return
	}

	fmt.Println(encryptedAESKey)
	aesKey, err := rsa.DecryptPKCS1v15(rand.Reader, pk, encryptedAESKey)
	fmt.Println("AESKey after decrypt", aesKey)
	if err != nil {
		fmt.Println("Erro ao descriptografar chave AES:", err)
		return
	}
	fmt.Println(aesKey)
}

//func DecryptWithAES () {
//	// Read the encrypted message from the connection
//	encryptedMessageBase64 := make([]byte, 2048)
//	if err != nil {
//	fmt.Println("Error reading encrypted message:", err)
//	return
//	}
//
//	encryptedMessage, err := base64.StdEncoding.DecodeString(string(encryptedMessageBase64[:n]))
//	if err != nil {
//	fmt.Println("Error decoding encrypted message:", err)
//	return
//	}
//
//	// Decrypt the message using the AES key
//	nonceSize := 12
//	if len(encryptedMessage) < nonceSize {
//	fmt.Println("Invalid encrypted message")
//	return
//	}
//
//	nonce, ciphertext := encryptedMessage[:nonceSize], encryptedMessage[nonceSize:]
//	block, err := aes.NewCipher(aesKey)
//	if err != nil {
//	fmt.Println("Error creating AES cipher:", err)
//	return
//	}
//
//	aesGCM, err := cipher.NewGCM(block)
//	if err != nil {
//	fmt.Println("Error creating GCM:", err)
//	return
//	}
//
//	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
//	if err != nil {
//	fmt.Println("Error decrypting message:", err)
//	return
//	}
//
//	fmt.Println("Decrypted message:", string(plaintext))
//}
