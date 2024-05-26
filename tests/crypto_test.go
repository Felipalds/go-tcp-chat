package tests

import (
	"fmt"
	"go-tcp-chat/encrypt"
	"testing"
)

func TestEncryption(t *testing.T) {

	// Ensure the key is 16, 24, or 32 bytes long

	key, err := encrypt.GenerateAESKey()
	if err != nil {
		t.Error("Error generating aes key", err)
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		t.Error("Key length must be 16, 24, or 32 bytes")
	}

	plainText := "Hello, World! adsfas fasdf asdf asdf asdf\n"

	encryptedText, err := encrypt.Encrypt(plainText, key)
	if err != nil {
		t.Error("Encryption failed ", err)
	}

	decryptedText, err := encrypt.Decrypt(encryptedText, key)
	if err != nil {
		t.Error("Decryption failed ", err)
	}

	if decryptedText != plainText {
		t.Error("Not equals!")
	}

	fmt.Println("Decrypted text:", decryptedText)
}
