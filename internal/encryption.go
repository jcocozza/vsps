package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

const (
	saltSize   = 16
	keySize    = 32
	iterations = 100000
)

// Derive a key from the user password
func deriveKey(password, salt []byte) []byte {
	key := append(password, salt...)
	for i := 0; i < iterations; i++ {
		h := hmac.New(sha256.New, key)
		h.Write(key)
		key = h.Sum(nil)
	}
	return key[:keySize]
}

// give a random salt
func generateSalt() []byte {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	return salt
}

// Encrypt data by key
func encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Main encrypt
func Encryptor(masterpass string, data []byte) ([]byte, error) {
	salt := generateSalt()
	key := deriveKey([]byte(masterpass), salt)
	ciphertext, err := encrypt(data, key)
	if err != nil {
		return nil, err
	}

	result := append(salt, ciphertext...)
	return result, nil
}

// decrypt data by key
func decrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("failed to decrypt: ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// Main decrypt
func Decryptor(masterpass string, data []byte) ([]byte, error) {
	// no data means nothing to decrypt
	if len(data) == 0 {
		return []byte{}, nil
	}

	if len(data) < saltSize {
		return nil, fmt.Errorf("ciphertext too short: less then salt length. data has been corrupted.")
	}

	salt := data[:saltSize]
	ciphertext := data[saltSize:]
	key := deriveKey([]byte(masterpass), salt)
	return decrypt(ciphertext, key)
}
