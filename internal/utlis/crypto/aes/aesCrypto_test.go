package aes

import (
	"testing"
)

func TestAESEncryptDecrypt(t *testing.T) {
	key := "secretkey"
	plainText := "Hello, Golang!"

	aesCrypto := &AES{}

	t.Run("EncryptDecryptValidInput", func(t *testing.T) {
		// Encrypt
		cipherText, err := aesCrypto.Encrypt(key, plainText)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		// Decrypt
		decryptedText, err := aesCrypto.Decrypt(key, cipherText)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		// Check if the decrypted text matches the original plain text
		if decryptedText != plainText {
			t.Fatalf("Decrypted text does not match original plain text. Got: %s, Expected: %s", decryptedText, plainText)
		}
	})

	t.Run("DecryptInvalidCipherText", func(t *testing.T) {
		invalidCipherText := "invalidciphertext"

		_, err := aesCrypto.Decrypt(key, invalidCipherText)

		// Check if decryption returns an error for invalid cipher text
		if err == nil {
			t.Fatalf("Expected an error for decryption with invalid cipher text, but got nil")
		}
	})

	t.Run("EncryptEmptyKey", func(t *testing.T) {
		emptyKey := ""
		_, err := aesCrypto.Encrypt(emptyKey, plainText)

		// Check if encryption returns an error for an empty key
		if err == nil {
			t.Fatalf("Expected an error for encryption with an empty key, but got nil")
		}
	})

	t.Run("DecryptEmptyKey", func(t *testing.T) {
		emptyKey := ""
		_, err := aesCrypto.Decrypt(emptyKey, "someciphertext")

		// Check if decryption returns an error for an empty key
		if err == nil {
			t.Fatalf("Expected an error for decryption with an empty key, but got nil")
		}
	})

	t.Run("DecryptShortCipherText", func(t *testing.T) {
		shortCipherText := "abc123"

		_, err := aesCrypto.Decrypt(key, shortCipherText)

		// Check if decryption returns an error for a short cipher text
		if err == nil {
			t.Fatalf("Expected an error for decryption with a short cipher text, but got nil")
		}
	})
}
