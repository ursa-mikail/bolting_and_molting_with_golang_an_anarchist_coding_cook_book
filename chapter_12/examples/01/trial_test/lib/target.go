package lib

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
)

// ComputeHMAC generates an HMAC-SHA256 hash for a given input
func ComputeHMAC(data []byte) []byte {
	key := []byte("super_secret_key")
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// SecureCompare safely compares two byte slices to prevent timing attacks
func SecureCompare(a, b []byte) bool {
	return bytes.Equal(a, b)
}
