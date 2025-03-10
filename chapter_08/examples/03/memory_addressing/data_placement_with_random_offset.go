package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Allocate a byte slice (memory region)
	mem := make([]byte, 1024) // 1 KB memory block

	// Pick a random offset within the allocated memory
	offset, _ := rand.Int(rand.Reader, big.NewInt(int64(len(mem))))
	randOffset := offset.Int64()

	// Generate a random byte
	var b [1]byte
	_, err := rand.Read(b[:])
	if err != nil {
		fmt.Println("Error generating random byte:", err)
		return
	}

	// Place the random byte at the chosen offset
	mem[randOffset] = b[0]

	// Print results
	fmt.Printf("Placed random byte %x at offset %d in allocated memory.\n", b[0], randOffset)
	fmt.Printf("Memory region content: %x\n", mem)
}

/*
Allocate memory, pick a random offset (random memory region address) within that allocated region, and modify the bytes (places random bytes there).

Placed random byte 45 at offset 611 in allocated memory.
Memory region content: 00000000...0045000000000...000000000000000000000000000
*/