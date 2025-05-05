// Demonstrates a custom serializer (streaming a string to bytes) and deserializer (rebuilding the ASCII string from bytes)
package main

import (
	"bytes"
	"fmt"
	"io"
)

// CustomSerDe provides serialization and deserialization
type CustomSerDe struct{}

// Serialize converts a string into a byte stream (io.Reader)
func (s *CustomSerDe) Serialize(input string) io.Reader {
	return bytes.NewReader([]byte(input))
}

// Deserialize reads from a byte stream and returns a string
func (s *CustomSerDe) Deserialize(r io.Reader) (string, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	serde := &CustomSerDe{}
	original := "Hello, 世界! ASCII + Unicode."

	// Serialize the string to a byte stream
	stream := serde.Serialize(original)

	// Deserialize back to string
	decoded, err := serde.Deserialize(stream)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original:  ", original)
	fmt.Println("Decoded:   ", decoded)
}

/*
bytes.NewReader: turns the byte slice into an io.Reader stream.
buf.ReadFrom: efficiently reads all bytes from the stream and reconstructs the string.

sample out:
Original:   Hello, 世界! ASCII + Unicode.
Decoded:    Hello, 世界! ASCII + Unicode.

*/