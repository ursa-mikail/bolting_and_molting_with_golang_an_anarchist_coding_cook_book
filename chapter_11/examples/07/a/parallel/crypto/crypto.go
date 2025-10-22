package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

type CryptoService struct {
	key []byte
}

func NewCryptoService(key []byte) *CryptoService {
	return &CryptoService{key: key}
}

func (c *CryptoService) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	// Create a GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	
	// Return base64 encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (c *CryptoService) Decrypt(encryptedData string) (string, error) {
	// Decode base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// VerifyRecovery checks if decrypted data matches original
func (c *CryptoService) VerifyRecovery(original, encryptedData string) (bool, error) {
	decrypted, err := c.Decrypt(encryptedData)
	if err != nil {
		return false, err
	}
	return decrypted == original, nil
}

// DemoParallelEncryption demonstrates parallel encryption with sync.WaitGroup and error channel
func (c *CryptoService) DemoParallelEncryption(data []string) ([]string, []error) {
	var wg sync.WaitGroup
	errCh := make(chan error, len(data))
	resultCh := make(chan struct {
		index int
		value string
	}, len(data))

	// Process each string in parallel
	for i, text := range data {
		wg.Add(1)
		go func(idx int, plaintext string) {
			defer wg.Done()

			// Simulate some processing time
			time.Sleep(time.Millisecond * 100)

			encrypted, err := c.Encrypt(plaintext)
			if err != nil {
				errCh <- fmt.Errorf("failed to encrypt item %d: %w", idx, err)
				return
			}

			// Verify recovery for demonstration
			recovered, err := c.VerifyRecovery(plaintext, encrypted)
			if err != nil || !recovered {
				errCh <- fmt.Errorf("recovery verification failed for item %d: %w", idx, err)
				return
			}

			resultCh <- struct {
				index int
				value string
			}{idx, encrypted}
		}(i, text)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errCh)
	close(resultCh)

	// Collect results
	results := make([]string, len(data))
	errors := make([]error, 0)

	// Read from error channel
	for err := range errCh {
		errors = append(errors, err)
	}

	// Read from result channel
	for result := range resultCh {
		results[result.index] = result.value
	}

	return results, errors
}

// DemoParallelWithWorkerPool demonstrates using a worker pool pattern
func (c *CryptoService) DemoParallelWithWorkerPool(data []string, numWorkers int) ([]string, []error) {
	var wg sync.WaitGroup
	jobs := make(chan struct {
		index int
		text  string
	}, len(data))
	results := make(chan struct {
		index int
		value string
		err   error
	}, len(data))

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				encrypted, err := c.Encrypt(job.text)
				if err != nil {
					results <- struct {
						index int
						value string
						err   error
					}{job.index, "", fmt.Errorf("worker %d: %w", workerID, err)}
					continue
				}

				// Verify recovery
				if recovered, err := c.VerifyRecovery(job.text, encrypted); err != nil || !recovered {
					results <- struct {
						index int
						value string
						err   error
					}{job.index, "", fmt.Errorf("worker %d recovery failed: %w", workerID, err)}
					continue
				}

				results <- struct {
					index int
					value string
					err   error
				}{job.index, encrypted, nil}
			}
		}(i)
	}

	// Send jobs
	for i, text := range data {
		jobs <- struct {
			index int
			text  string
		}{i, text}
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	finalResults := make([]string, len(data))
	var finalErrors []error

	for result := range results {
		if result.err != nil {
			finalErrors = append(finalErrors, result.err)
		} else {
			finalResults[result.index] = result.value
		}
	}

	return finalResults, finalErrors
}
