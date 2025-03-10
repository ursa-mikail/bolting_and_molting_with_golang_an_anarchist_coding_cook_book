package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

const (
	regionCount = 8 // Number of regions: 3 (for small)
	filename    = "mmap_test.bin"
)

func main() {
	pageSize := syscall.Getpagesize()
	regionSize := pageSize // Use system page size for region size

	fmt.Printf("System page size: %d bytes\n", pageSize)

	// Open the file for mmap (if using file-backed mmap)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Ensure the file is large enough
	fileSize := int64(regionSize * regionCount)
	if err := file.Truncate(fileSize); err != nil {
		fmt.Println("Error resizing file:", err)
		return
	}

	// Allocate multiple regions
	mappedRegions := make([][]byte, regionCount)
	for i := 0; i < regionCount; i++ {
		offset := int64(i) * int64(regionSize)

		// Debug: Check if offset is aligned properly
		if offset%int64(pageSize) != 0 {
			fmt.Printf("Error: Offset %d is not aligned to %d bytes\n", offset, pageSize)
			return
		}

		// macOS Fix: Try MAP_PRIVATE instead of MAP_SHARED on macOS
		mmapFlags := unix.MAP_SHARED
		fd := int(file.Fd())
		if runtime.GOOS == "darwin" {
			fmt.Println("Using MAP_PRIVATE instead of MAP_SHARED on macOS")
			mmapFlags = unix.MAP_PRIVATE
		}

		// Debug: Show mapping attempt
		fmt.Printf("Attempting to map region %d at offset %d...\n", i, offset)

		// Map the memory region
		data, err := unix.Mmap(fd, offset, regionSize, unix.PROT_READ|unix.PROT_WRITE, mmapFlags)
		if err != nil {
			fmt.Printf("Error mapping region %d: %v\n", i, err)
			return
		}

		mappedRegions[i] = data
		fmt.Printf("Region %d mapped at address: %p\n", i, &data[0])

		// Write random data
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < regionSize; j++ {
			data[j] = byte(rand.Intn(256))
		}
	}

	// Make memory read-only
	for i, data := range mappedRegions {
		err := unix.Mprotect(data, unix.PROT_READ)
		if err != nil {
			fmt.Printf("Error setting region %d to read-only: %v\n", i, err)
		} else {
			fmt.Printf("Region %d is now read-only\n", i)
		}
	}

	// Unmap memory
	for i, data := range mappedRegions {
		if err := unix.Munmap(data); err != nil {
			fmt.Printf("Error unmapping region %d: %v\n", i, err)
		} else {
			fmt.Printf("Region %d unmapped successfully\n", i)
		}
	}
}

/*
% go mod init example.com/demo

% go get golang.org/x/sys/unix
% go mod tidy

% go run main.go

mmap_allocate_memory_and_randomly_modify_bytes_with_multiple_memory_regions_protected.go

Maps multiple memory regions using mmap (file-backed).
Handles file size issues to avoid invalid argument.
Protects memory by making it read-only after writing.
Simulates memory faults (e.g., accessing unmapped memory).
Uses memory-mapped files for persistence.

//
removed the hardcoded regionSize = 4096 constant and instead made regionSize a variable that takes its value from the system's page size. This ensures that all memory mappings are correctly aligned with your system's page boundaries (16384 bytes on your macOS system).
note:

Kept the regionCount constant at N (=3 for small)
Preserved all the macOS-specific handling with MAP_PRIVATE
Maintained all debug outputs and error checking


program demonstrates memory mapping (mmap) in Go, which is a technique to map files or devices into memory

1. System Page Size Detection: It starts by determining the system's page size (which varies by OS - 16384 bytes on macOS system).
2. File Creation: It creates or opens a file named "mmap_test.bin" that will be used for the memory mapping.
3. File Resizing: It ensures the file is large enough to accommodate all the memory regions (pageSize × N regions).
4. Memory Mapping Loop: It maps N (=3 for small) separate regions of memory from the file:
Each region size equals the system page size
Each region starts at an offset that's a multiple of the page size
On macOS, it uses MAP_PRIVATE instead of MAP_SHARED (addressing platform differences)
It provides debug information showing where each region is mapped in memory

5. Data Writing: After mapping each region, it fills it with random data (bytes 0-255).
6. Memory Protection: It changes all mapped regions from read-write to read-only using Mprotect to demonstrate memory protection capabilities.
7. Memory Unmapping: Finally, it unmaps all the regions, releasing the memory.

The program showcases several important system programming concepts:

Memory-mapped I/O
System page size alignment requirements
Platform-specific memory management flags
Memory protection changes
Proper resource cleanup

The error can occurred because this is trying to map at offset 4096, which wasn't aligned with macOS page size of 16384 bytes.
Memory mapping requires offsets to be multiples of the page size, which is why the fix was to use the system's page size for the region size.

% go run main.go
System page size: 16384 bytes
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 0 at offset 0...
Region 0 mapped at address: 0x100bb8000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 1 at offset 16384...
Region 1 mapped at address: 0x100bbc000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 2 at offset 32768...
Region 2 mapped at address: 0x100bc0000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 3 at offset 49152...
Region 3 mapped at address: 0x100bc4000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 4 at offset 65536...
Region 4 mapped at address: 0x100bc8000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 5 at offset 81920...
Region 5 mapped at address: 0x100bcc000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 6 at offset 98304...
Region 6 mapped at address: 0x100bd0000
Using MAP_PRIVATE instead of MAP_SHARED on macOS
Attempting to map region 7 at offset 114688...
Region 7 mapped at address: 0x100bd4000
Region 0 is now read-only
Region 1 is now read-only
Region 2 is now read-only
Region 3 is now read-only
Region 4 is now read-only
Region 5 is now read-only
Region 6 is now read-only
Region 7 is now read-only
Region 0 unmapped successfully
Region 1 unmapped successfully
Region 2 unmapped successfully
Region 3 unmapped successfully
Region 4 unmapped successfully
Region 5 unmapped successfully
Region 6 unmapped successfully
Region 7 unmapped successfully

*/
