package jwd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type JWTHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	Encrypted bool   `json:"enc"` // Indicates payload is encrypted
}

// GenerateHMACSignature creates a secure HMAC-SHA256 signature
func GenerateHMACSignature(data []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// VerifyHMACSignature verifies the HMAC-SHA256 signature
func VerifyHMACSignature(data []byte, key []byte, signature []byte) bool {
	expectedSignature := GenerateHMACSignature(data, key)
	return hmac.Equal(expectedSignature, signature)
}

// GenerateNonce creates a cryptographically secure random nonce
func GenerateNonce() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Encrypt encrypts data using AES-256-GCM
func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate nonce for GCM
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Encrypt and authenticate
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	// Prepend nonce to ciphertext
	return append(nonce, ciphertext...), nil
}

// Decrypt decrypts data using AES-256-GCM
func Decrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) < 12 {
		return nil, errors.New("invalid ciphertext")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Split nonce and ciphertext
	nonce, ciphertext := data[:12], data[12:]

	// Decrypt and verify
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
