package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yourusername/grpc-example/proto"
	"github.com/yourusername/grpc-example/crypto"
)

type server struct {
	pb.UnimplementedSecureServiceServer
	cryptoService *crypto.CryptoService
}

func (s *server) ProcessData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	log.Printf("Received data: %s", req.Data)

	// Encrypt the data
	encryptedData, err := s.cryptoService.Encrypt(req.Data)
	if err != nil {
		return &pb.DataResponse{
			Success: false,
		}, err
	}

	// Verify recovery (for demonstration)
	recovered, err := s.cryptoService.VerifyRecovery(req.Data, encryptedData)
	if err != nil {
		log.Printf("Recovery verification failed: %v", err)
	} else {
		log.Printf("Data recovery verified: %t", recovered)
	}

	return &pb.DataResponse{
		EncryptedData: encryptedData,
		Success:       true,
	}, nil
}

func main() {
	// Use a 32-byte key for AES-256
	key := []byte("thisis32bitlongpassphraseimusing")
	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	cryptoService := crypto.NewCryptoService(key)
	pb.RegisterSecureServiceServer(s, &server{
		cryptoService: cryptoService,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
