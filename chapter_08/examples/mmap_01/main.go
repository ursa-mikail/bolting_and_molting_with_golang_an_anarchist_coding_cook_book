package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Constants for mmap
const (
	ArraySize   = 1024 * 1024 // 1MB of floats (262,144 floats)
	ElementSize = 8           // Size of float64 in bytes
	MapSize     = ArraySize * ElementSize
)

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create a random array
	fmt.Println("=== Step 1: Creating Random Array ===")
	originalArray := createRandomArray(ArraySize)
	fmt.Printf("Created array with %d elements\n", len(originalArray))
	fmt.Printf("First 5 elements: %v\n", originalArray[:5])

	// Create and use memory-mapped file
	fmt.Println("\n=== Step 2: Using Memory-Mapped File ===")
	err := useMemoryMappedFile(originalArray)
	if err != nil {
		log.Fatal("Error using memory-mapped file:", err)
	}

	fmt.Println("\n=== Demo Completed Successfully ===")
}

func createRandomArray(size int) []float64 {
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Float64() * 1000
	}
	return arr
}

func useMemoryMappedFile(original []float64) error {
	// Create a temporary file for mmap
	tempFile, err := os.CreateTemp("", "mmap-demo-*.dat")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up temp file
	defer tempFile.Close()

	// Set file size to match our array size
	if err := tempFile.Truncate(MapSize); err != nil {
		return fmt.Errorf("failed to set file size: %w", err)
	}

	// Memory map the file
	fmt.Println("Memory mapping file...")
	mappedData, err := unix.Mmap(int(tempFile.Fd()), 0, MapSize,
		unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return fmt.Errorf("mmap failed: %w", err)
	}
	defer unix.Munmap(mappedData) // Always unmap when done

	// Convert mapped memory to float64 slice for easy access
	mappedArray := (*[ArraySize]float64)(unsafe.Pointer(&mappedData[0]))
	mappedSlice := mappedArray[:]

	// Copy original array to mapped memory
	fmt.Println("Copying data to mapped memory...")
	copy(mappedSlice, original)

	// Verify the copy
	fmt.Println("\n=== Step 3: Verifying Copy ===")
	fmt.Printf("First 5 elements in mapped memory: %v\n", mappedSlice[:5])

	// Calculate and compare sums to verify data integrity
	originalSum := sumArray(original)
	mappedSum := sumArray(mappedSlice)

	fmt.Printf("Original array sum: %.2f\n", originalSum)
	fmt.Printf("Mapped array sum: %.2f\n", mappedSum)
	fmt.Printf("Sums match: %v\n", almostEqual(originalSum, mappedSum))

	// Modify the mapped memory
	fmt.Println("\n=== Step 4: Modifying Mapped Memory ===")
	fmt.Println("Doubling all values in mapped memory...")
	for i := range mappedSlice {
		mappedSlice[i] *= 2
	}
	fmt.Printf("First 5 elements after modification: %v\n", mappedSlice[:5])

	// Clear mapped memory (set all to zero)
	fmt.Println("\n=== Step 5: Clearing Mapped Memory ===")
	clearMemory(mappedData)

	// Verify all elements are zero
	fmt.Println("Verifying all elements are zero...")
	allZero := true
	for i := 0; i < 10; i++ { // Check first 10 elements
		if mappedSlice[i] != 0 {
			allZero = false
			break
		}
	}
	fmt.Printf("First 10 elements are zero: %v\n", allZero)

	// Create a second mapping to show persistence
	fmt.Println("\n=== Step 6: Demonstrating Persistence ===")
	secondMap(tempFile.Name())

	return nil
}

func sumArray(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-9
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}

func clearMemory(data []byte) {
	fmt.Printf("Clearing %d bytes of memory...\n", len(data))
	for i := range data {
		data[i] = 0
	}
	fmt.Println("Memory cleared successfully")
}

func secondMap(filename string) {
	// Open existing file
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Error opening file for second mapping: %v", err)
		return
	}
	defer file.Close()

	// Map the file again
	mappedData, err := unix.Mmap(int(file.Fd()), 0, MapSize,
		unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		log.Printf("Second mmap failed: %v", err)
		return
	}
	defer unix.Munmap(mappedData)

	// Access as float64 array
	mappedArray := (*[ArraySize]float64)(unsafe.Pointer(&mappedData[0]))
	mappedSlice := mappedArray[:]

	// Check first few values
	fmt.Printf("First 5 elements in second mapping: %v\n", mappedSlice[:5])
	fmt.Println("Note: All values should be zero after clearing")
}

/*
Memory Mapping: The unix.Mmap() function maps the file into memory
Type Conversion: Uses unsafe.Pointer to interpret bytes as float64 array
Memory Safety: Always defers Munmap() to ensure cleanup
Direct Memory Access: Modifications to mappedSlice directly affect file contents
Clearing Memory: Shows how to zero out all bytes in mapped memory
Persistence: Demonstrates that mapped memory persists in the file

=== Step 1: Creating Random Array ===
Created array with 1048576 elements
First 5 elements: [834.523420047032 158.79580554129768 238.5991641825794 102.87405901711068 57.83374037489826]

=== Step 2: Using Memory-Mapped File ===
Memory mapping file...
Copying data to mapped memory...

=== Step 3: Verifying Copy ===
First 5 elements in mapped memory: [834.523420047032 158.79580554129768 238.5991641825794 102.87405901711068 57.83374037489826]
Original array sum: 524692469.17
Mapped array sum: 524692469.17
Sums match: true

=== Step 4: Modifying Mapped Memory ===
Doubling all values in mapped memory...
First 5 elements after modification: [1669.046840094064 317.59161108259536 477.1983283651588 205.74811803422136 115.66748074979652]

=== Step 5: Clearing Mapped Memory ===
Clearing 8388608 bytes of memory...
Memory cleared successfully
Verifying all elements are zero...
First 10 elements are zero: true

=== Step 6: Demonstrating Persistence ===
First 5 elements in second mapping: [0 0 0 0 0]
Note: All values should be zero after clearing

=== Demo Completed Successfully ===
*/
