/*
Written to a file (simulating serialization to a file stream),
Then read back (simulating deserialization),
Converted back to the original ASCII/Unicode string.
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type FileSerDe struct{}

// SerializeToFile writes a string to a file as bytes
func (f *FileSerDe) SerializeToFile(input string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bytes.NewReader([]byte(input))
	_, err = io.Copy(file, reader)
	return err
}

// DeserializeFromFile reads from a file and reconstructs the string
func (f *FileSerDe) DeserializeFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func main() {
	serde := &FileSerDe{}
	original := "Stream to file and back! 文件流 🎉"
	filename := "output_stream.txt"

	// Serialize to file
	err := serde.SerializeToFile(original, filename)
	if err != nil {
		panic(err)
	}

	// Deserialize from file
	decoded, err := serde.DeserializeFromFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original:  ", original)
	fmt.Println("Decoded:   ", decoded)
}


/*
SerializeToFile uses os.Create and io.Copy to write the string into a file via a stream.
DeserializeFromFile reads the file as a byte stream and reconstructs the string.

This mimics how many messaging systems or logging tools might persist payloads to disk and then read them back.

sample out:
Original:   Stream to file and back! 文件流 🎉
Decoded:    Stream to file and back! 文件流 🎉

*/