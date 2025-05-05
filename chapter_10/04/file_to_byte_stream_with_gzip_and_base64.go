/*
Compresses the input string using GZIP,
Encodes the compressed data using Base64,
Writes the Base64 string to a file,
Reads the file back,
Decodes the Base64,
Decompresses the GZIP,
And reconstructs the original string.
*/

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

type AdvancedFileSerDe struct{}

// CompressAndEncode compresses the input using GZIP and encodes it in Base64
func (s *AdvancedFileSerDe) CompressAndEncode(input string) (string, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(input))
	if err != nil {
		return "", err
	}
	gz.Close()

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encoded, nil
}

// DecodeAndDecompress decodes Base64 and decompresses GZIP to recover original string
func (s *AdvancedFileSerDe) DecodeAndDecompress(encoded string) (string, error) {
	compressedData, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	r, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return "", err
	}
	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// SerializeToFile writes compressed+encoded string to file
func (s *AdvancedFileSerDe) SerializeToFile(input, filename string) error {
	encoded, err := s.CompressAndEncode(input)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(encoded), 0644)
}

// DeserializeFromFile reads, decodes and decompresses to recover original string
func (s *AdvancedFileSerDe) DeserializeFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return s.DecodeAndDecompress(string(data))
}

func main() {
	serde := &AdvancedFileSerDe{}
	original := "This is a test string with 🎉 unicode and ASCII. 文件流 + gzip + base64"
	filename := "compressed_encoded.txt"

	// Serialize
	err := serde.SerializeToFile(original, filename)
	if err != nil {
		panic(err)
	}

	// Deserialize
	decoded, err := serde.DeserializeFromFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original:  ", original)
	fmt.Println("Decoded:   ", decoded)
}

/*
Uses compress/gzip to reduce size and obscure content.
Uses encoding/base64 to safely store binary data as text in a file.
The file remains portable and human-readable (Base64 text).

sample out:
Original:   This is a test string with 🎉 unicode and ASCII. 文件流 + gzip + base64
Decoded:    This is a test string with 🎉 unicode and ASCII. 文件流 + gzip + base64
*/