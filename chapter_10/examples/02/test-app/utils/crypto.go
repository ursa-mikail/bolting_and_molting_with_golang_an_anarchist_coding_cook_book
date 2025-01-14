package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// GenerateHash creates a SHA-256 hash of the given data.
func GenerateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// GenerateHMAC creates an HMAC using the provided secret key and data.
func GenerateHMAC(secretKey, data string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}
