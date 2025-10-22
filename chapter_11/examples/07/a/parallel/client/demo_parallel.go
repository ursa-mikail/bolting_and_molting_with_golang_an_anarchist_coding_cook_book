// Demo Client Showing Parallel Execution
package main

import (
	"log"
	"time"

	"github.com/yourusername/grpc-example/crypto"
)

func demoParallelEncryption() {
	key := []byte("thisis32bitlongpassphraseimusing")
	cryptoService := crypto.NewCryptoService(key)

	// Test data for parallel processing
	testData := []string{
		"Hello World 1",
		"Hello World 2", 
		"Hello World 3",
		"Hello World 4",
		"Hello World 5",
		"", // Empty string to test error handling
		"Hello World 7",
	}

	log.Println("=== Demo: Basic Parallel Encryption ===")
	start := time.Now()
	
	results, errors := cryptoService.DemoParallelEncryption(testData)
	
	elapsed := time.Since(start)
	log.Printf("Parallel encryption completed in %v", elapsed)

	// Print results
	for i, result := range results {
		if result != "" {
			log.Printf("Item %d: Encrypted to %d chars", i, len(result))
		}
	}

	// Print errors
	if len(errors) > 0 {
		log.Printf("Encountered %d errors:", len(errors))
		for _, err := range errors {
			log.Printf("  - %v", err)
		}
	}

	log.Println("\n=== Demo: Worker Pool Pattern ===")
	start = time.Now()
	
	workerResults, workerErrors := cryptoService.DemoParallelWithWorkerPool(testData, 3)
	
	elapsed = time.Since(start)
	log.Printf("Worker pool encryption completed in %v", elapsed)

	// Verify some results were recovered correctly
	successCount := 0
	for i, result := range workerResults {
		if result != "" {
			successCount++
			// Verify we can decrypt it
			decrypted, err := cryptoService.Decrypt(result)
			if err == nil && decrypted == testData[i] {
				log.Printf("✓ Item %d successfully encrypted and recovered", i)
			}
		}
	}
	
	log.Printf("Successfully processed: %d/%d items", successCount, len(testData))
	if len(workerErrors) > 0 {
		log.Printf("Worker pool errors: %d", len(workerErrors))
	}
}

func main() {
	demoParallelEncryption()
}
