package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random address within a certain range (for demonstration, using a high heap address)
	randomAddress := uintptr(rand.Int63n(0x7fffffffffff))

	// Generate a random byte
	randomByte := byte(rand.Intn(256))

	// Convert the address to a pointer
	ptr := (*byte)(unsafe.Pointer(randomAddress))

	// Attempt to write to the random memory address (DANGEROUS)
	fmt.Printf("Attempting to write byte 0x%x to address 0x%x\n", randomByte, randomAddress)

	// Try writing to the address (this can crash the program)
	syscall.Mprotect(unsafe.Slice(ptr, 1), syscall.PROT_READ|syscall.PROT_WRITE)
	*ptr = randomByte

	fmt.Println("Memory modified successfully!")
}
