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
		
		log.Printf("Chunk %d: Encrypting %d bytes", i, len(chunk))
		encrypted, err := fp.CryptoService.EncryptChunk(chunk)
		if err != nil {
			log.Printf("Chunk %d: Encryption failed: %v", i, err)
			failed++
			chunkResult.Error = fmt.Sprintf("encrypt failed: %v", err)
			chunkResult.Success = false
			chunkResults[i] = chunkResult
			continue
		}
		log.Printf("Chunk %d: Encrypted to %d bytes", i, len(encrypted))

		log.Printf("Chunk %d: Decrypting...", i)
		decrypted, err := fp.CryptoService.DecryptChunk(encrypted)
		if err != nil {
			log.Printf("Chunk %d: Decryption failed: %v", i, err)
			failed++
			chunkResult.Error = fmt.Sprintf("decrypt failed: %v", err)
			chunkResult.Success = false
			chunkResults[i] = chunkResult
			continue
		}
		log.Printf("Chunk %d: Decrypted to %d bytes", i, len(decrypted))

		if !bytes.Equal(chunk, decrypted) {
			log.Printf("Chunk %d: Integrity check failed (original: %d, decrypted: %d)", i, len(chunk), len(decrypted))
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
		log.Printf("Chunk %d: SUCCESS", i)
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