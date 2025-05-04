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
	"fmt"
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
	// Load certs
	cert, err := tls.LoadX509KeyPair(folder+"client.crt", folder+"client.key")
	if err != nil {
		log.Fatal("failed loading client cert/key:", err)
	}

	caCert, err := os.ReadFile(folder + "ca.crt")
	if err != nil {
		log.Fatal("failed reading CA cert:", err)
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caPool,
		InsecureSkipVerify: false,
		ServerName:         "localhost", // must match SAN in server cert
	}

	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	// ECDH key pair
	priv, x, y, _ := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	pub := PublicKey{X: x.Bytes(), Y: y.Bytes()}
	_ = enc.Encode(pub)

	var serverPub PublicKey
	_ = dec.Decode(&serverPub)

	sx := new(big.Int).SetBytes(serverPub.X)
	sy := new(big.Int).SetBytes(serverPub.Y)
	sharedX, _ := elliptic.P256().ScalarMult(sx, sy, priv)
	sharedKey := sha256.Sum256(sharedX.Bytes())

	block, _ := aes.NewCipher(sharedKey[:])
	aesgcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesgcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	// Send nonce
	conn.Write(nonce)

	// Send encrypted file
	data, _ := os.ReadFile("file_to_send")
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)
	conn.Write(ciphertext)

	// Receive and compare hash
	recvHash := make([]byte, 32)
	conn.Read(recvHash)
	hash := sha256.Sum256(data)

	if string(hash[:]) == string(recvHash) {
		fmt.Println("✅ File verified!")
	} else {
		fmt.Println("❌ Hash mismatch!")
	}
}

/*
% go run client.go
✅ File verified!
*/
