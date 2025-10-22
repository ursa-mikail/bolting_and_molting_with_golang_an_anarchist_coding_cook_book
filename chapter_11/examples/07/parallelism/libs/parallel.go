package libs

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"

	"parallelism/crypto"
	"parallelism/types"
)

func (fp *FileProcessor) ProcessParallel(chunks [][]byte, originalHash string) (*types.ProcessingResult, time.Duration, error) {
	log.Println("Starting PARALLEL processing...")
	start := time.Now()

	var wg sync.WaitGroup
	errCh := make(chan error, len(chunks))

	decryptedChunks := make([][]byte, len(chunks))
	chunkResults := make([]types.ChunkResult, len(chunks))
	var successCount, failCount int
	var mu sync.Mutex

	for i, chunk := range chunks {
		wg.Add(1)
		go func(idx int, chunkData []byte) {
			defer wg.Done()

			chunkResult := types.ChunkResult{Index: idx}
			
			encrypted, err := fp.CryptoService.EncryptChunk(chunkData)
			if err != nil {
				mu.Lock()
				failCount++
				mu.Unlock()
				chunkResult.Error = fmt.Sprintf("encrypt failed: %v", err)
				chunkResult.Success = false
				chunkResults[idx] = chunkResult
				errCh <- fmt.Errorf("chunk %d: %w", idx, err)
				return
			}

			decrypted, err := fp.CryptoService.DecryptChunk(encrypted)
			if err != nil {
				mu.Lock()
				failCount++
				mu.Unlock()
				chunkResult.Error = fmt.Sprintf("decrypt failed: %v", err)
				chunkResult.Success = false
				chunkResults[idx] = chunkResult
				errCh <- fmt.Errorf("chunk %d: %w", idx, err)
				return
			}

			if !bytes.Equal(chunkData, decrypted) {
				mu.Lock()
				failCount++
				mu.Unlock()
				chunkResult.Error = "integrity check failed"
				chunkResult.Success = false
				chunkResults[idx] = chunkResult
				errCh <- fmt.Errorf("chunk %d integrity failed", idx)
				return
			}

			decryptedChunks[idx] = decrypted
			chunkResult.Success = true
			chunkResults[idx] = chunkResult
			
			mu.Lock()
			successCount++
			mu.Unlock()
		}(i, chunk)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			log.Printf("Error in goroutine: %v", err)
		}
	}

	elapsed := time.Since(start)
	log.Printf("Parallel processing completed in %v", elapsed)

	reconstructed := fp.ReconstructData(decryptedChunks)
	integrityVerified := VerifyIntegrity(originalHash, reconstructed)

	result := &types.ProcessingResult{
		Chunks:            chunkResults,
		OriginalHash:      originalHash,
		ReconstructedHash: crypto.ComputeSHA256(reconstructed),
		IntegrityVerified: integrityVerified,
		Stats: types.ProcessingStats{
			TotalChunks:      len(chunks),
			SuccessfulChunks: successCount,
			FailedChunks:     failCount,
			ProcessingTime:   elapsed,
			SpeedMBPerSecond: CalculateSpeed(len(reconstructed), elapsed),
		},
	}

	return result, elapsed, nil
}