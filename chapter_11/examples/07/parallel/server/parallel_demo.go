// Server with Parallel Demo Endpoint
package main

import (
	"context"
	"log"

	pb "github.com/yourusername/grpc-example/proto"
	"github.com/yourusername/grpc-example/crypto"
)

// Add this method to the server struct
func (s *server) ProcessBatchData(ctx context.Context, req *pb.BatchDataRequest) (*pb.BatchDataResponse, error) {
	log.Printf("Processing batch of %d items", len(req.Data))

	// Use parallel encryption
	results, errors := s.cryptoService.DemoParallelEncryption(req.Data)

	// Convert to response
	response := &pb.BatchDataResponse{
		EncryptedData: results,
		SuccessCount:  int32(len(results) - len(errors)),
		ErrorCount:    int32(len(errors)),
	}

	if len(errors) > 0 {
		errorMessages := make([]string, len(errors))
		for i, err := range errors {
			errorMessages[i] = err.Error()
		}
		response.Errors = errorMessages
	}

	return response, nil
}