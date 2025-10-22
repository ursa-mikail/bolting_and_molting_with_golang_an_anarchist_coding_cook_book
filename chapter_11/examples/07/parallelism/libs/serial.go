package libs

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"parallelism/crypto"
	"parallelism/types"
)

func (fp *FileProcessor) ProcessSerial(chunks [][]byte, originalHash string) (*types.ProcessingResult, time.Duration, error) {
	log.Println("Starting SERIAL processing...")
	start := time.Now()

	decryptedChunks := make([][]byte, len(chunks))
	chunkResults := make([]types.ChunkResult, len(chunks))
	successful := 0
	failed := 0

	for i, chunk := range chunks {
		chunkResult := types.ChunkResult{Index: i}
		
		encrypted, err := fp.CryptoService.EncryptChunk(chunk)
		if err != nil {
			failed++
			chunkResult.Error = fmt.Sprintf("encrypt failed: %v", err)
			chunkResult.Success = false
			chunkResults[i] = chunkResult
			continue
		}

		decrypted, err := fp.CryptoService.DecryptChunk(encrypted)
		if err != nil {
			failed++
			chunkResult.Error = fmt.Sprintf("decrypt failed: %v", err)
			chunkResult.Success = false
			chunkResults[i] = chunkResult
			continue
		}

		if !bytes.Equal(chunk, decrypted) {
			failed++
			chunkResult.Error = "integrity check failed"
			chunkResult.Success = false
			chunkResults[i] = chunkResult
			continue
		}

		decryptedChunks[i] = decrypted
		chunkResult.Success = true
		chunkResults[i] = chunkResult
		successful++
	}

	elapsed := time.Since(start)
	log.Printf("Serial processing completed in %v", elapsed)

	reconstructed := fp.ReconstructData(decryptedChunks)
	integrityVerified := VerifyIntegrity(originalHash, reconstructed)

	result := &types.ProcessingResult{
		Chunks:            chunkResults,
		OriginalHash:      originalHash,
		ReconstructedHash: crypto.ComputeSHA256(reconstructed),
		IntegrityVerified: integrityVerified,
		Stats: types.ProcessingStats{
			TotalChunks:      len(chunks),
			SuccessfulChunks: successful,
			FailedChunks:     failed,
			ProcessingTime:   elapsed,
			SpeedMBPerSecond: CalculateSpeed(len(reconstructed), elapsed),
		},
	}

	return result, elapsed, nil
}