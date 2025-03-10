package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"unsafe"
)

func main() {
	N := 3 * 1024 * 1024 // 3MB
	M := 1024            // Distance between addresses (1KB)
	K := 16              // Number of random bytes to fill

	// Allocate N bytes
	mem := make([]byte, N)

	// Get a random offset ensuring room for M
	maxOffset := int64(N - M)
	if maxOffset <= 0 {
		fmt.Println("Invalid M value, must be smaller than N")
		return
	}

	offset1Big, err := rand.Int(rand.Reader, big.NewInt(maxOffset))
	if err != nil {
		panic(err)
	}
	offset1 := offset1Big.Int64()
	offset2 := offset1 + int64(M)

	// Ensure K bytes can fit at both locations
	if offset1+int64(K) > int64(N) || offset2+int64(K) > int64(N) {
		fmt.Println("K value is too large for selected addresses")
		return
	}

	// Fill both locations with random bytes
	rand.Read(mem[offset1 : offset1+int64(K)])
	rand.Read(mem[offset2 : offset2+int64(K)])

	// Print memory addresses and values
	fmt.Printf("Address 1: %p, Data: %x\n", unsafe.Pointer(&mem[offset1]), mem[offset1:offset1+int64(K)])
	fmt.Printf("Address 2: %p, Data: %x\n", unsafe.Pointer(&mem[offset2]), mem[offset2:offset2+int64(K)])
}

/*
1. allocates N bytes using make([]byte, N).
2. selects a random offset1 within bounds so that offset2 = offset1 + M remains valid.
3. ensures K bytes fit at both locations.
4. fills K bytes at both addresses using crypto/rand.Read().
5. prints the addresses and their corresponding random hex values.

Converted offset1 from *big.Int to int64 using .Int64()
Ensured correct type handling for offset1 + int64(K) to avoid mismatches
Ensured maxOffset is properly checked before choosing offset1


Address 1: 0xc0002725d1, Data: 739c129859aa68ce4fbb8470e65148e8
Address 2: 0xc0002729d1, Data: 6e0c18aeaf7cc61efb9e8acfffcc6b18
*/