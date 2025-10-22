package types

import "time"

type ProcessingStats struct {
	TotalChunks          int
	SuccessfulChunks     int
	FailedChunks         int
	ProcessingTime       time.Duration
	SpeedMBPerSecond     float64
}

type ChunkResult struct {
	Index          int
	Data           []byte
	EncryptedData  []byte
	DecryptedData  []byte
	Error          string
	Success        bool
}

type ProcessingResult struct {
	Chunks             []ChunkResult
	OriginalHash       string
	ReconstructedHash  string
	IntegrityVerified  bool
	Stats              ProcessingStats
}