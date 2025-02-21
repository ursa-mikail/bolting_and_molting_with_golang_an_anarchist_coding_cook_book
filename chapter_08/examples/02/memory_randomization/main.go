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

/*
Sample run:
% go run main.go
Original string: you should be able to see me
Hex representation: 796f752073686f756c642062652061626c6520746f20736565206d65
Before deletion, memory contains (string): you should be able to see me
After randomization, memory contains (string): !
                                                ??
                                                  ?_r??2???
                                                           FO?
                                                              ?`??+??J
After randomization, memory contains (in hex): 210c8c940bf05f72f3fa32f8bdf60c464fa20bdb6082df1f2bfbd44a
Reference deleted. Original content should no longer be retrievable.
*/
