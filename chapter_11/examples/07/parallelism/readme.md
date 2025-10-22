go mod init parallelism
go mod tidy

### Generate the Go code:
```bash
protoc --go_out=. --go_opt=paths=source_relative proto/types.proto
```

### Run the application
go run main.go libs/common.go libs/serial.go libs/parallel.go crypto/crypto.go proto/types.pb.go

### Or build and run
go build -o parallelism .
./parallelism

```
Protocol Buffers: Structured data definitions for processing results

Modular Libraries:

common.go: Shared utilities and base functionality

serial.go: Sequential processing implementation

parallel.go: Concurrent processing with goroutines
```

### make

```
# Initialize the project properly
make init

# Download dependencies
make deps

# Check everything is set up correctly
make check

# Build and run
make run
```

```
Development:

bash
make fmt && make lint && make test && make run-dev

Production build:

bash
make clean && make build && make test

```

```
% go run test_crypto.go
Testing crypto service:
Key length: 35
Test data: Hello, World!
Encrypted: 41 bytes
Decrypted: Hello, World!
✓ Crypto test PASSED
```

```
% make run
Building parallelism... 
go build -o parallelism main.go
Running parallelism... 
./parallelism
2025/10/21 19:25:29 === FILE PROCESSOR BENCHMARK ===
2025/10/21 19:25:29 Creating 100MB test zip file...
2025/10/21 19:25:31 Created zip: 104889745 bytes, SHA256: c903fb7a5b9fa6fd0c18d8699d2413a358648e0a0eaa26cd9b4c32251d4cb3a7
2025/10/21 19:25:31 Split into 51 chunks
2025/10/21 19:25:31 
============================================================
2025/10/21 19:25:31 SERIAL PROCESSING (Sequential For Loop)
2025/10/21 19:25:31 ============================================================
2025/10/21 19:25:31 Starting SERIAL processing...
2025/10/21 19:25:31 Serial processing completed in 12.708µs
2025/10/21 19:25:31 
============================================================
2025/10/21 19:25:31 PARALLEL PROCESSING (Goroutines with Sync)
2025/10/21 19:25:31 ============================================================
2025/10/21 19:25:31 Starting PARALLEL processing...
2025/10/21 19:25:31 Error in goroutine: chunk 50: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 25: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 26: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 27: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 28: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 29: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 30: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 31: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 32: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 33: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 34: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 35: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 36: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 37: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 38: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 39: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 40: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 41: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 42: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 43: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 44: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 45: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 46: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 24: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 47: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 11: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 48: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 49: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 12: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 5: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 0: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 13: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 1: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 14: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 6: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 2: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 15: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 3: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 7: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 16: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 4: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 17: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 8: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 18: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 9: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 10: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 19: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 21: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 22: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 20: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Error in goroutine: chunk 23: crypto/aes: invalid key size 35
2025/10/21 19:25:31 Parallel processing completed in 370.708µs
2025/10/21 19:25:31 
============================================================
2025/10/21 19:25:31 RESULTS COMPARISON
2025/10/21 19:25:31 ============================================================
2025/10/21 19:25:31 SERIAL PROCESSING:
2025/10/21 19:25:31   ✓ Integrity Verified: false
2025/10/21 19:25:31   ✓ Chunks: 0 successful, 51 failed
2025/10/21 19:25:31   ⏱️  Pure Processing Time: 12.708µs
2025/10/21 19:25:31   ⏱️  Total Time (with setup): 19.625µs
2025/10/21 19:25:31   🚀 Processing Speed: 0.00 MB/s
2025/10/21 19:25:31   🚀 Total Speed: 5097103.48 MB/s
2025/10/21 19:25:31 
PARALLEL PROCESSING:
2025/10/21 19:25:31   ✓ Integrity Verified: false
2025/10/21 19:25:31   ✓ Chunks: 0 successful, 51 failed
2025/10/21 19:25:31   ⏱️  Pure Processing Time: 370.708µs
2025/10/21 19:25:31   ⏱️  Total Time (with setup): 378.292µs
2025/10/21 19:25:31   🚀 Processing Speed: 0.00 MB/s
2025/10/21 19:25:31   🚀 Total Speed: 264427.10 MB/s
2025/10/21 19:25:31 
🏆 PERFORMANCE IMPROVEMENT:
2025/10/21 19:25:31   📉 Time Reduction: -2817.12%
2025/10/21 19:25:31   📈 Speed Increase: NaN%
2025/10/21 19:25:31   ⚡ Efficiency Factor: 0.03x
2025/10/21 19:25:31 
============================================================
2025/10/21 19:25:31 PARALLELISM ANALYSIS
2025/10/21 19:25:31 ============================================================
2025/10/21 19:25:31 WHERE PARALLELISM OCCURS:
2025/10/21 19:25:31   • Goroutine Creation: 51 chunks → 51 concurrent goroutines
2025/10/21 19:25:31   • Concurrent Encryption: All 51 chunks encrypted simultaneously
2025/10/21 19:25:31   • Concurrent Decryption: All 51 chunks decrypted simultaneously
2025/10/21 19:25:31   • Sync Mechanisms:
2025/10/21 19:25:31     - sync.WaitGroup: Coordinates completion of all 51 goroutines
2025/10/21 19:25:31     - Mutex: Protects shared counters (success/failure counts)
2025/10/21 19:25:31     - Buffered Error Channel: Non-blocking error collection (102 capacity)
2025/10/21 19:25:31   • Go Runtime Scheduling:
2025/10/21 19:25:31     - Goroutines distributed across available CPU cores
2025/10/21 19:25:31     - Automatic load balancing by Go scheduler
2025/10/21 19:25:31     - Efficient context switching between goroutines
2025/10/21 19:25:31 
RESULTS VALIDATION:
2025/10/21 19:25:31   • SHA256 Match: false (Serial: e3b0c44298fc1c14..., Parallel: e3b0c44298fc1c14...)
2025/10/21 19:25:31   • Chunk Success Rate: Serial 0/51, Parallel 0/51
```


