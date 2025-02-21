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
*/
