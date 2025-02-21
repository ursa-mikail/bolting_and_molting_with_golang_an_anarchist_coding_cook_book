package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// RandomizeMemory fills a byte slice with random data
func RandomizeMemory(b []byte) {
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = byte(rand.Intn(256)) // Random byte value (0-255)
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

	// Print before deletion
	fmt.Println("Before deletion, memory contains (string):", string(data))

	// Randomize memory
	RandomizeMemory(data)
	fmt.Println("After randomization, memory contains (string):", string(data))
	fmt.Println("After randomization, memory contains (in hex):", hex.EncodeToString(data))

	// "Delete" reference by setting data to nil
	data = nil

	// Force garbage collection (optional, to speed up memory cleanup)
	// runtime.GC()

	// The original data should not be accessible anymore
	fmt.Println("Reference deleted. Original content should no longer be retrievable.")
}
