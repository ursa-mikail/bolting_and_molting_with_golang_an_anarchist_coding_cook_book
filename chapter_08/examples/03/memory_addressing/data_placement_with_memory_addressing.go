package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Allocate an array of bytes
	mem := make([]byte, 1024) // 1 KB

	// Get the base address of the allocated slice
	baseAddr := uintptr(unsafe.Pointer(&mem[0]))

	// Pick a random offset within the allocated region
	randomOffset := uintptr(rand.Intn(len(mem)))

	// Calculate the random memory address
	randomAddr := baseAddr + randomOffset

	// Access memory at that address
	ptr := (*byte)(unsafe.Pointer(randomAddr))
	*ptr = byte(rand.Intn(256)) // Place a random byte

	fmt.Printf("Allocated memory starts at: 0x%x\n", baseAddr)
	fmt.Printf("Randomly chosen address: 0x%x\n", randomAddr)
	fmt.Printf("Value placed at address 0x%x: 0x%x\n", randomAddr, *ptr)
}

/*
Choosing memory addresses in Go is tricky because Go runs in a managed environment with garbage collection (GC), memory safety, and runtime protections that prevent direct manipulation of arbitrary memory addresses.

1. Allocate memory with make(), ensuring Go manages it.
2. Only modify memory within the allocated range.


Allocated memory starts at: 0xc000104ae8
Randomly chosen address: 0xc000104b73
Value placed at address 0xc000104b73: 0x9f
*/