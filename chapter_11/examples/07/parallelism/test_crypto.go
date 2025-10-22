package main

import (
	"fmt"
	"log"

	"parallelism/crypto"
)

func main() {
	key := []byte("thisis32bitlongpassphraseimusing123")
	service := crypto.NewCryptoService(key)
	
	testData := []byte("Hello, World!")
	
	fmt.Printf("Testing crypto service:\n")
	fmt.Printf("Key length: %d\n", len(key))
	fmt.Printf("Test data: %s\n", string(testData))
	
	encrypted, err := service.EncryptChunk(testData)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}
	fmt.Printf("Encrypted: %d bytes\n", len(encrypted))
	
	decrypted, err := service.DecryptChunk(encrypted)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %s\n", string(decrypted))
	
	if string(testData) == string(decrypted) {
		fmt.Println("✓ Crypto test PASSED")
	} else {
		fmt.Println("✗ Crypto test FAILED")
	}
}