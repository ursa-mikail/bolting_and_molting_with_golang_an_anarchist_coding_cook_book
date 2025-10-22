package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/yourusername/grpc-example/proto"
	"github.com/yourusername/grpc-example/crypto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSecureServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test data
	testData := "Hello, secure gRPC world!"

	r, err := client.ProcessData(ctx, &pb.DataRequest{Data: testData})
	if err != nil {
		log.Fatalf("could not process data: %v", err)
	}

	log.Printf("Encrypted data: %s", r.GetEncryptedData())
	log.Printf("Success: %t", r.GetSuccess())

	// Verify we can decrypt it on client side too
	key := []byte("thisis32bitlongpassphraseimusing")
	cryptoService := crypto.NewCryptoService(key)

	// Verify recovery
	recovered, err := cryptoService.VerifyRecovery(testData, r.GetEncryptedData())
	if err != nil {
		log.Fatalf("Recovery verification failed: %v", err)
	}

	log.Printf("Data recovery verified on client: %t", recovered)
}
