package libs

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"log"
	"time"

	"parallelism/crypto"
)

const (
	ChunkSize = 2 * 1024 * 1024
	FileSize  = 100 * 1024 * 1024
)

type FileProcessor struct {
	CryptoService *crypto.CryptoService
}

func NewFileProcessor(key []byte) *FileProcessor {
	return &FileProcessor{
		CryptoService: crypto.NewCryptoService(key),
	}
}

func (fp *FileProcessor) CreateTestZip() ([]byte, string, error) {
	log.Println("Creating 100MB test zip file...")
	testData := make([]byte, FileSize)
	if _, err := rand.Read(testData); err != nil {
		return nil, "", err
	}

	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuffer)
	fileWriter, err := zipWriter.Create("test_data.bin")
	if err != nil {
		return nil, "", err
	}
	if _, err := fileWriter.Write(testData); err != nil {
		return nil, "", err
	}
	if err := zipWriter.Close(); err != nil {
		return nil, "", err
	}

	zipData := zipBuffer.Bytes()
	originalHash := crypto.ComputeSHA256(zipData)
	log.Printf("Created zip: %d bytes, SHA256: %s\n", len(zipData), originalHash)
	return zipData, originalHash, nil
}

func (fp *FileProcessor) SplitIntoChunks(data []byte) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i += ChunkSize {
		end := i + ChunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	log.Printf("Split into %d chunks", len(chunks))
	return chunks
}

func (fp *FileProcessor) ReconstructData(chunks [][]byte) []byte {
	var totalSize int
	for _, chunk := range chunks {
		totalSize += len(chunk)
	}
	reconstructed := make([]byte, 0, totalSize)
	for _, chunk := range chunks {
		reconstructed = append(reconstructed, chunk...)
	}
	return reconstructed
}

func CalculateSpeed(dataSize int, duration time.Duration) float64 {
	sizeMB := float64(dataSize) / (1024 * 1024)
	seconds := duration.Seconds()
	if seconds == 0 {
		return 0
	}
	return sizeMB / seconds
}

func VerifyIntegrity(originalHash string, reconstructed []byte) bool {
	return originalHash == crypto.ComputeSHA256(reconstructed)
}