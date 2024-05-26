package encrypt

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

func DecryptAESKey(key string, pk *rsa.PrivateKey) []byte {
	key = strings.Replace(key, "\n", "", -1)
	encryptedAESKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Println("Erro ao decodificar chave AES criptografada:", err)
		return nil
	}

	aesKey, err := rsa.DecryptPKCS1v15(rand.Reader, pk, encryptedAESKey)
	if err != nil {
		fmt.Println("Erro ao descriptografar chave AES:", err)
		return nil
	}
	return aesKey
}
