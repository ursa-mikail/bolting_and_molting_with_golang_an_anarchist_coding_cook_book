package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random memory address (not guaranteed to be safe)
	randomAddress := uintptr(rand.Int63n(0x7fffffffffff)) // Targeting a high memory region

	// Convert to pointer
	ptr := (*byte)(unsafe.Pointer(randomAddress))

	// Attempt to write to the random memory address
	fmt.Printf("Attempting to write to memory address: 0x%x\n", randomAddress)
	*ptr = byte(rand.Intn(256)) // This will likely crash

	fmt.Println("Successfully wrote to memory!") // This line likely won't be reached
}

/*
⚠ DANGER! This approach can:

Crash the program with a segmentation fault.
Access protected memory, leading to unpredictable behavior.
Violate OS security restrictions.
*/
