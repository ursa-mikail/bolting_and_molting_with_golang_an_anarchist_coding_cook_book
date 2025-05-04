package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/gob"
	"io"
	"log"
	"math/big"
	"os"
)

type PublicKey struct {
	X, Y []byte
}

func main() {
	folder := "./certs/"

	cert, err := tls.LoadX509KeyPair(folder+"server.crt", folder+"server.key")
	if err != nil {
		log.Fatalf("Error loading server cert/key: %v", err)
	}

	caCert, err := os.ReadFile(folder + "ca.crt")
	if err != nil {
		log.Fatalf("Error reading CA cert: %v", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caPool,
	}

	ln, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatalf("Server failed to listen: %v", err)
	}
	defer func() {
		log.Println("Server shutting down gracefully.")
		ln.Close()
	}()

	log.Println("✅ Server listening on port 8443...")

	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer func() {
		log.Println("Closing connection with client.")
		conn.Close()
	}()

	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	// ECDH key exchange
	priv, x, y, _ := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	pub := PublicKey{X: x.Bytes(), Y: y.Bytes()}
	if err := enc.Encode(pub); err != nil {
		log.Fatalf("Error sending public key: %v", err)
	}

	var clientPub PublicKey
	if err := dec.Decode(&clientPub); err != nil {
		log.Fatalf("Error receiving client public key: %v", err)
	}

	cx := new(big.Int).SetBytes(clientPub.X)
	cy := new(big.Int).SetBytes(clientPub.Y)
	sharedX, _ := elliptic.P256().ScalarMult(cx, cy, priv)
	sharedKey := sha256.Sum256(sharedX.Bytes())

	block, _ := aes.NewCipher(sharedKey[:])
	aesgcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(conn, nonce); err != nil {
		log.Fatalf("Failed to read nonce from client: %v", err)
	}

	// Read encrypted file data
	encryptedData := make([]byte, 4096)
	n, err := conn.Read(encryptedData)
	if err != nil {
		log.Fatalf("Failed to read encrypted data: %v", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, encryptedData[:n], nil)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	err = os.WriteFile("received_file", plaintext, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
	log.Println("✅ Received and saved file.")

	// Compute hash and send back
	hash := sha256.Sum256(plaintext)
	_, err = conn.Write(hash[:])
	if err != nil {
		log.Fatalf("Failed to send back hash: %v", err)
	}

	log.Println("✅ Hash sent to client for verification.")
}

/*
% go run server.go
2025/05/01 10:30:18 ✅ Server listening on port 8443...
2025/05/01 10:30:21 ✅ Received and saved file.
2025/05/01 10:30:21 ✅ Hash sent to client for verification.
2025/05/01 10:30:21 Closing connection with client.
2025/05/01 10:30:21 Server shutting down gracefully.
*/
