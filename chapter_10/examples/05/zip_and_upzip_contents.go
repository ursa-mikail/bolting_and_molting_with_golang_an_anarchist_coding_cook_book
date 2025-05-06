/*
This demonstrates validating data integrity across serialization boundaries, such as:
1. Ensuring files haven't been tampered with or corrupted after archival.
2. Testing backup and restore pipelines or verifying signed bundles.
3. Verifying behavior in CI/CD workflows, secure storage, or air-gapped environments.
4. Developing something like a secure file transfer system or digital forensics tooling.

The use of SHA256, timestamps, and folder structure implies you're being precise — not just checking contents, but possibly testing reproducibility or tamper detection.

Creates a folder.
Creates files in it with content and custom timestamps.
Computes SHA256 hashes for each file.
Zips the folder.
Unzips it into another folder.
Recomputes SHA256s and compares them.
*/

package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Step 1: Create original folder and files
	origDir := "original_folder"
	os.MkdirAll(origDir, 0755)
	files := map[string]string{
		"file1.txt": "Hello from file one",
		"file2.txt": "Another file with content",
	}

	hashes := make(map[string]string)

	for name, content := range files {
		path := filepath.Join(origDir, name)
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			panic(err)
		}
		// Set custom timestamp
		timestamp := time.Date(2022, time.January, 1, 10, 0, 0, 0, time.UTC)
		os.Chtimes(path, timestamp, timestamp)

		hash, err := hashFile(path)
		if err != nil {
			panic(err)
		}
		hashes[name] = hash
		fmt.Printf("Original hash of %s: %s\n", name, hash)
	}

	// Step 2: Zip the folder
	zipPath := "archive.zip"
	err := zipFolder(origDir, zipPath)
	if err != nil {
		panic(err)
	}

	// Step 3: Unzip into new folder
	destDir := "unzipped_folder"
	os.MkdirAll(destDir, 0755)
	err = unzip(zipPath, destDir)
	if err != nil {
		panic(err)
	}

	// Step 4: Verify SHA256
	for name, origHash := range hashes {
		newPath := filepath.Join(destDir, origDir, name)
		newHash, err := hashFile(newPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Unzipped hash of %s: %s\n", name, newHash)
		if newHash != origHash {
			fmt.Printf("❌ Mismatch for %s\n", name)
		} else {
			fmt.Printf("✅ Match for %s\n", name)
		}
	}
}

// Compute SHA256 of a file
func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Zip a folder
func zipFolder(folder, destZip string) error {
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(filepath.Dir(folder), path)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		fh.Name = relPath
		fh.Method = zip.Deflate

		writer, err := archive.CreateHeader(fh)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		return err
	})
	return err
}

// Unzip a zip file into destination
func unzip(srcZip, destDir string) error {
	r, err := zip.OpenReader(srcZip)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		srcFile, err := f.Open()
		if err != nil {
			dstFile.Close()
			return err
		}

		_, err = io.Copy(dstFile, srcFile)
		dstFile.Close()
		srcFile.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

/*
Original hash of file1.txt: ee26d4eb923699cd7fe5cb22e81b455c8d57f4722c463d82d3fe0b882bb56691
Original hash of file2.txt: a098a3a96e830eb66b013380fb8ebc2b8fd563226378185432624eb70d98ea53
Unzipped hash of file1.txt: ee26d4eb923699cd7fe5cb22e81b455c8d57f4722c463d82d3fe0b882bb56691
✅ Match for file1.txt
Unzipped hash of file2.txt: a098a3a96e830eb66b013380fb8ebc2b8fd563226378185432624eb70d98ea53
✅ Match for file2.txt
*/