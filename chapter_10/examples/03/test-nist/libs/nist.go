package libs

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// NISTTest represents the top-level JSON structure
type NISTTest struct {
	Algorithm string     `json:"algorithm"`
	Mode      string     `json:"algoMode"`
	Tests     TestGroups `json:"testGroups"`
}

// TestGroups represents a group of test vectors
type TestGroups []struct {
	Direction string `json:"direction"`
	KeyLen    int    `json:"keyLen"`
	IVLen     int    `json:"ivLen"`
	PTLen     int    `json:"payloadLen"`
	AADLen    int    `json:"aadLen"`
	TagLen    int    `json:"tagLen"`
	Tests     []struct {
		ID       int    `json:"tcId"`
		Key      string `json:"key"`
		IV       string `json:"iv"`
		PT       string `json:"pt"`
		AAD      string `json:"aad"`
		CT       string `json:"ct"`
		Tag      string `json:"tag"`
		TestPass bool   `json:"testPassed"`
	} `json:"tests"`
}

func FetchNISTTest(url string) (*NISTTest, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var test NISTTest
	if err := json.Unmarshal(body, &test); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &test, nil
}

func RunTest(test *NISTTest) {
	fmt.Printf("Running tests for Algorithm: %s\nMode: %s\n\n", test.Algorithm, test.Mode)

	for groupIdx, group := range test.Tests {
		fmt.Printf("Test Group %d (%s):\n", groupIdx+1, group.Direction)
		fmt.Printf("Key Length: %d, IV Length: %d, Tag Length: %d\n\n", 
			group.KeyLen, group.IVLen, group.TagLen)

		for _, t := range group.Tests {
			fmt.Printf("Test Case %d:\n", t.ID)

			key, err := hex.DecodeString(t.Key)
			if err != nil {
				log.Printf("Error decoding key: %v", err)
				continue
			}

			iv, err := hex.DecodeString(t.IV)
			if err != nil {
				log.Printf("Error decoding IV: %v", err)
				continue
			}

			pt, err := hex.DecodeString(t.PT)
			if err != nil {
				log.Printf("Error decoding plaintext: %v", err)
				continue
			}

			aad, err := hex.DecodeString(t.AAD)
			if err != nil {
				log.Printf("Error decoding AAD: %v", err)
				continue
			}

			if group.Direction == "encrypt" {
				// Create a custom GCM with the specified tag length
				ct, tag, err := EncryptAESGCMWithParams(key, iv, pt, aad, group.TagLen/8)
				if err != nil {
					log.Printf("Encryption failed: %v", err)
					continue
				}

				expectedCT, err := hex.DecodeString(t.CT)
				if err != nil {
					log.Printf("Error decoding expected ciphertext: %v", err)
					continue
				}

				expectedTag, err := hex.DecodeString(t.Tag)
				if err != nil {
					log.Printf("Error decoding expected tag: %v", err)
					continue
				}

				if !compareBytes(ct, expectedCT) || !compareBytes(tag, expectedTag) {
					fmt.Printf("❌ Failed: Mismatch in ciphertext or tag\n")
					fmt.Printf("Expected CT: %X\nGot CT:      %X\n", expectedCT, ct)
					fmt.Printf("Expected Tag: %X\nGot Tag:      %X\n", expectedTag, tag)
					continue
				}
				fmt.Printf("✓ Passed\n")

			} else { // decrypt
				ct, err := hex.DecodeString(t.CT)
				if err != nil {
					log.Printf("Error decoding ciphertext: %v", err)
					continue
				}

				tag, err := hex.DecodeString(t.Tag)
				if err != nil {
					log.Printf("Error decoding tag: %v", err)
					continue
				}

				decryptedPT, err := DecryptAESGCMWithParams(key, iv, ct, tag, aad, group.TagLen/8)
				if err != nil {
					fmt.Printf("❌ Failed: Decryption error - %v\n", err)
					continue
				}

				if !compareBytes(decryptedPT, pt) {
					fmt.Printf("❌ Failed: Decrypted plaintext mismatch\n")
					fmt.Printf("Expected PT: %X\nGot PT:      %X\n", pt, decryptedPT)
					continue
				}
				fmt.Printf("✓ Passed\n")
			}
		}
		fmt.Println()
	}
}


func EncryptAESGCMWithParams(key, iv, plaintext, aad []byte, tagLen int) ([]byte, []byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to create cipher: %w", err)
    }

    // Use the full IV length provided in the test case
    nonce := iv // Use the full 120-bit IV as provided

    gcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
    if err != nil {
        return nil, nil, fmt.Errorf("failed to create GCM: %w", err)
    }

    // Encrypt and get the tag
    sealed := gcm.Seal(nil, nonce, plaintext, aad)
    ciphertext := sealed[:len(sealed)-16] // GCM always produces a 16-byte tag
    tag := sealed[len(sealed)-16:]

    // Truncate tag if necessary
    if tagLen < len(tag) {
        tag = tag[:tagLen]
    }

    return ciphertext, tag, nil
}


func DecryptAESGCMWithParams(key, iv, ciphertext, tag, aad []byte, tagLen int) ([]byte, error) {
    // Check if inputs are valid
    if len(key) == 0 || len(iv) == 0 || len(tag) == 0 {
        return nil, fmt.Errorf("key, IV, or tag cannot be empty")
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, fmt.Errorf("failed to create cipher: %w", err)
    }

    // Use the full IV length as provided
    nonce := iv

    // Create a GCM cipher, ensure the nonce size matches the length of iv
    gcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
    if err != nil {
        return nil, fmt.Errorf("failed to create GCM: %w", err)
    }

    // Ensure tag length matches the expected size for GCM
    if tagLen <= 0 || tagLen > len(tag) {
        return nil, fmt.Errorf("invalid tag length")
    }

    // Combine ciphertext and tag correctly, appending the correct tag length
    ciphertextAndTag := append(ciphertext, tag[:tagLen]...)

    // Perform decryption and authentication
    plaintext, err := gcm.Open(nil, nonce, ciphertextAndTag, aad)
    if err != nil {
        fmt.Printf("Decryption Error: %v\n", err) // Add more detailed error logs
        return nil, fmt.Errorf("decryption error: %w", err)
    }

    return plaintext, nil
}


func compareBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
