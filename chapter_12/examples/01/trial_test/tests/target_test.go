package tests

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"testing"
)

// TestDefensePass checks if two identical byte slices produce the same HMAC
func TestDefensePass(t *testing.T) {
	secret1 := []byte("random_secret_value_1")
	secret2 := secret1

	hmac1, _ := computeHMAC(secret1)
	hmac2, _ := computeHMAC(secret2)

	if !bytes.Equal(hmac1, hmac2) {
		t.Errorf("Defense failed: HMACs do not match")
	}
}

// TestDefenseFail checks if different byte slices produce different HMACs
func TestDefenseFail(t *testing.T) {
	secret1 := []byte("random_secret_value_1")
	secret2 := []byte("random_secret_value_2")

	hmac1, _ := computeHMAC(secret1)
	hmac2, _ := computeHMAC(secret2)

	if bytes.Equal(hmac1, hmac2) {
		t.Errorf("Defense failed: HMACs should not match")
	}
}

// BenchmarkCacheDefense measures performance of byte comparison
func BenchmarkDefense(b *testing.B) {
	secret1 := []byte("benchmark_secret_value_1")
	secret2 := secret1

	b.ResetTimer()
	for N := 16; N <= 2048*2; N *= 2 {
		b.Run(fmt.Sprintf("Size_%d", N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bytes.Equal(secret1, secret2)
			}
		})
	}
}

// computeHMAC generates an HMAC-SHA256 hash for a given input
func computeHMAC(data []byte) ([]byte, error) {
	if data == nil {
		return nil, fmt.Errorf("input data cannot be nil")
	}

	key := []byte("super_secret_key")
	h := hmac.New(sha256.New, key)
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}



/*

trial_test % go test -v ./tests
trial_test %
Run only TestDefense[Pass/Fail]:
go test -v -run TestDefensePass ./tests
go test -v -run TestDefenseFail ./tests

Run only BenchmarkDefense:
go test -bench BenchmarkDefense ./tests

Running Specific Tests
Run all tests (both functional and benchmark tests):
go test -v ./tests


Run both functional and benchmark tests together explicitly:
go test -v -run 'TestDefensePass|TestDefenseFail' -bench BenchmarkCacheDefense ./tests

Run all benchmark tests (ignore normal tests):
go test -bench . ./tests

% go test ./tests
ok      example.com/demo/tests  0.381s

% go test -bench=. ./tests 
goos: darwin
goarch: arm64
pkg: example.com/demo/tests
cpu: Apple M1
BenchmarkDefense/Size_16-8          506158704            2.068 ns/op
BenchmarkDefense/Size_32-8          576874540            2.083 ns/op
BenchmarkDefense/Size_64-8          579767796            2.072 ns/op
BenchmarkDefense/Size_128-8         576997975            2.085 ns/op
BenchmarkDefense/Size_256-8         574993299            2.068 ns/op
BenchmarkDefense/Size_512-8         575291240            2.086 ns/op
BenchmarkDefense/Size_1024-8        579340242            2.067 ns/op
BenchmarkDefense/Size_2048-8        576988728            2.078 ns/op
BenchmarkDefense/Size_4096-8        580605243            2.096 ns/op
PASS
ok      example.com/demo/tests  12.809s

% go test -cover ./tests
ok      example.com/demo/tests  0.223s  coverage: [no statements]

*/
