package main

import (
	"log"
	"time"

	"parallelism/libs"  
	"parallelism/types"	
)

func main() {
	// 32-byte key for AES-256
	key := []byte("thisis32bitlongpassphraseimusing123")

	// Initialize processor
	fp := libs.NewFileProcessor(key)

	// Step 1: Create 100MB test zip file
	log.Println("=== FILE PROCESSOR BENCHMARK ===")
	zipData, originalHash, err := fp.CreateTestZip()
	if err != nil {
		log.Fatalf("Failed to create test zip: %v", err)
	}

	// Step 2: Split into 2MB chunks
	chunks := fp.SplitIntoChunks(zipData)

	// Step 3: Process chunks SERIALLY
	log.Println("\n" + "============================================================")
	log.Println("SERIAL PROCESSING (Sequential For Loop)")
	log.Println("============================================================")
	
	serialStart := time.Now()
	serialResult, serialDuration, err := fp.ProcessSerial(chunks, originalHash)
	if err != nil {
		log.Fatalf("Serial processing failed: %v", err)
	}
	serialTotalTime := time.Since(serialStart)

	// Step 4: Process chunks in PARALLEL
	log.Println("\n" + "============================================================")
	log.Println("PARALLEL PROCESSING (Goroutines with Sync)")
	log.Println("============================================================")
	
	parallelStart := time.Now()
	parallelResult, parallelDuration, err := fp.ProcessParallel(chunks, originalHash)
	if err != nil {
		log.Fatalf("Parallel processing failed: %v", err)
	}
	parallelTotalTime := time.Since(parallelStart)

	// Step 5: Display Results
	log.Println("\n" + "============================================================")
	log.Println("RESULTS COMPARISON")
	log.Println("============================================================")
	
	displayResults(serialResult, parallelResult, serialDuration, parallelDuration, 
		serialTotalTime, parallelTotalTime, len(zipData))

	// Step 6: Parallelism Insights
	log.Println("\n" + "============================================================")
	log.Println("PARALLELISM ANALYSIS")
	log.Println("============================================================")
	analyzeParallelism(chunks, serialResult, parallelResult)
}

func displayResults(serial, parallel *types.ProcessingResult, 
	serialDur, parallelDur, serialTotal, parallelTotal time.Duration, dataSize int) {
	
	log.Printf("SERIAL PROCESSING:")
	log.Printf("  ✓ Integrity Verified: %t", serial.IntegrityVerified)
	log.Printf("  ✓ Chunks: %d successful, %d failed", serial.Stats.SuccessfulChunks, serial.Stats.FailedChunks)
	log.Printf("  ⏱️  Pure Processing Time: %v", serialDur)
	log.Printf("  ⏱️  Total Time (with setup): %v", serialTotal)
	log.Printf("  🚀 Processing Speed: %.2f MB/s", serial.Stats.SpeedMBPerSecond)
	log.Printf("  🚀 Total Speed: %.2f MB/s", libs.CalculateSpeed(dataSize, serialTotal))

	log.Printf("\nPARALLEL PROCESSING:")
	log.Printf("  ✓ Integrity Verified: %t", parallel.IntegrityVerified)
	log.Printf("  ✓ Chunks: %d successful, %d failed", parallel.Stats.SuccessfulChunks, parallel.Stats.FailedChunks)
	log.Printf("  ⏱️  Pure Processing Time: %v", parallelDur)
	log.Printf("  ⏱️  Total Time (with setup): %v", parallelTotal)
	log.Printf("  🚀 Processing Speed: %.2f MB/s", parallel.Stats.SpeedMBPerSecond)
	log.Printf("  🚀 Total Speed: %.2f MB/s", libs.CalculateSpeed(dataSize, parallelTotal))

	// Calculate improvements
	timeReduction := (float64(serialDur) - float64(parallelDur)) / float64(serialDur) * 100
	speedImprovement := (parallel.Stats.SpeedMBPerSecond - serial.Stats.SpeedMBPerSecond) / serial.Stats.SpeedMBPerSecond * 100

	log.Printf("\n🏆 PERFORMANCE IMPROVEMENT:")
	log.Printf("  📉 Time Reduction: %.2f%%", timeReduction)
	log.Printf("  📈 Speed Increase: %.2f%%", speedImprovement)
	log.Printf("  ⚡ Efficiency Factor: %.2fx", float64(serialDur)/float64(parallelDur))
}

func analyzeParallelism(chunks [][]byte, serial, parallel *types.ProcessingResult) {
	log.Printf("WHERE PARALLELISM OCCURS:")
	log.Printf("  • Goroutine Creation: %d chunks → %d concurrent goroutines", len(chunks), len(chunks))
	log.Printf("  • Concurrent Encryption: All %d chunks encrypted simultaneously", len(chunks))
	log.Printf("  • Concurrent Decryption: All %d chunks decrypted simultaneously", len(chunks))
	log.Printf("  • Sync Mechanisms:")
	log.Printf("    - sync.WaitGroup: Coordinates completion of all %d goroutines", len(chunks))
	log.Printf("    - Mutex: Protects shared counters (success/failure counts)")
	log.Printf("    - Buffered Error Channel: Non-blocking error collection (%d capacity)", len(chunks)*2)
	log.Printf("  • Go Runtime Scheduling:")
	log.Printf("    - Goroutines distributed across available CPU cores")
	log.Printf("    - Automatic load balancing by Go scheduler")
	log.Printf("    - Efficient context switching between goroutines")
	
	log.Printf("\nRESULTS VALIDATION:")
	log.Printf("  • SHA256 Match: %t (Serial: %s, Parallel: %s)", 
		serial.IntegrityVerified && parallel.IntegrityVerified,
		serial.ReconstructedHash[:16] + "...",
		parallel.ReconstructedHash[:16] + "...")
	log.Printf("  • Chunk Success Rate: Serial %d/%d, Parallel %d/%d",
		serial.Stats.SuccessfulChunks, serial.Stats.TotalChunks,
		parallel.Stats.SuccessfulChunks, parallel.Stats.TotalChunks)
}
