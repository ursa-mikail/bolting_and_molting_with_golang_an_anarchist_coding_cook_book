package main

import (
	"fmt"

	"example.com/demo/lib"
)

func main() {
	secret1 := []byte("random_secret_value_1")
	//secret2 := []byte("random_secret_value_2")

	// Compute HMAC for both secrets
	hmac1 := lib.ComputeHMAC(secret1)
	hmac2 := lib.ComputeHMAC(secret1)

	// Compare the computed HMACs
	if lib.SecureCompare(hmac1, hmac2) {
		fmt.Println("Secrets match!")
	} else {
		fmt.Println("Secrets do NOT match!")
	}
}

/*
% go mod init example.com/demo
% go run main.go

*/
