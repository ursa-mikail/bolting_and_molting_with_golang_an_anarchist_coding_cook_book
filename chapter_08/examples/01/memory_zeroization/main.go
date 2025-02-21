package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

// ZeroMemory securely wipes a byte slice
func ZeroMemory(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

func main() {
	// Original string
	original := "you should be able to see me"
	fmt.Println("Original string:", original)

	// Convert to hex
	hexEncoded := hex.EncodeToString([]byte(original))
	fmt.Println("Hex representation:", hexEncoded)

	// Store in memory as a slice
	data := []byte(original)

	// Store unsafe pointer to raw memory
	ptr := (*string)(unsafe.Pointer(&data))

	// Print before deletion
	fmt.Println("Before deletion, raw memory read:", *ptr)

	// "Delete" reference by setting data to nil
	data = nil

	// Try to read the data back using the pointer
	fmt.Println("After deletion, raw memory read:", *ptr) // Should still print original content

	// Securely zero out memory before deletion
	realData := []byte(original)
	ZeroMemory(realData)

	// Attempt to read after zeroization
	fmt.Println("After zeroization, reading memory:", string(realData)) // Should be empty
}
