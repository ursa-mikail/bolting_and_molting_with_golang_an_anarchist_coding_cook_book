package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

func getSSHConfig(user, password, keyPath string) (*ssh.ClientConfig, error) {
	var authMethod ssh.AuthMethod

	if keyPath != "" {
		key, err := os.ReadFile(keyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read private key: %w", err)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		authMethod = ssh.PublicKeys(signer)
	} else {
		authMethod = ssh.Password(password)
	}

	return &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}

func createTestZip(fileName, zipName string) (string, string, error) {
	content := []byte("Hello from Go SSH automation!\n")
	err := ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		return "", "", err
	}

	zipFile, err := os.Create(zipName)
	if err != nil {
		return "", "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	fw, err := zipWriter.Create(filepath.Base(fileName))
	if err != nil {
		return "", "", err
	}
	_, err = fw.Write(content)
	if err != nil {
		return "", "", err
	}
	zipWriter.Close()

	hashBytes, err := os.ReadFile(zipName)
	if err != nil {
		return "", "", err
	}
	sum := sha256.Sum256(hashBytes)
	return zipName, hex.EncodeToString(sum[:]), nil
}

func uploadAndVerify(zipName, remotePath, remoteHash string, sshClient *ssh.Client) error {
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Open SCP-compatible stdin pipe
	writer, err := session.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer writer.Close()
		info, _ := os.Stat(zipName)
		fmt.Fprintf(writer, "C0644 %d %s\n", info.Size(), filepath.Base(remotePath))
		io.Copy(writer, mustOpen(zipName))
		fmt.Fprint(writer, "\x00")
	}()

	if err := session.Run(fmt.Sprintf("scp -t %s", filepath.Dir(remotePath))); err != nil {
		return fmt.Errorf("scp failed: %w", err)
	}

	// Now verify and unzip
	cmds := fmt.Sprintf(`
cd %s &&
actual=$(sha256sum %s | awk '{print $1}') &&
if [ "$actual" = "%s" ]; then unzip -o %s; else echo "SHA mismatch!"; exit 1; fi
`, filepath.Dir(remotePath), filepath.Base(remotePath), remoteHash, filepath.Base(remotePath))

	sess2, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer sess2.Close()

	var output bytes.Buffer
	sess2.Stdout = &output
	sess2.Stderr = &output

	err = sess2.Run(cmds)
	fmt.Println("Remote Output:\n", output.String())

	return err
}

func mustOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	host := os.Getenv("SSH_HOST")
	pass := os.Getenv("SSH_PASS")
	key := os.Getenv("SSH_KEY")

	if host == "" {
		log.Fatal("Missing SSH_HOST")
	}

	parts := strings.SplitN(host, "@", 2)
	if len(parts) != 2 {
		log.Fatalf("Invalid SSH_HOST: %s", host)
	}
	user, addr := parts[0], parts[1]

	config, err := getSSHConfig(user, pass, key)
	if err != nil {
		log.Fatal("SSH config error:", err)
	}

	client, err := ssh.Dial("tcp", addr+":22", config)
	if err != nil {
		log.Fatal("SSH dial failed:", err)
	}
	defer client.Close()

	zipName, localHash, err := createTestZip("hello.txt", "upload.zip")
	if err != nil {
		log.Fatal("Zip failed:", err)
	}

	fmt.Println("Local SHA256:", localHash)

	err = uploadAndVerify(zipName, "~/upload.zip", localHash, client)
	if err != nil {
		log.Fatal("Upload failed:", err)
	}
}

/* ssh_upload.go
// for SSH variables, refer: https://github.com/ursa-mikail/shell_script_utility/blob/main/scripts/utilities/dev_shell.sh
% go mod tidy
% go run main.go
Local SHA256: 749688b90fdadc9485b984c905ee747a212a8b40a4d520d930d65401dd21dbc1
Remote Output:
 Archive:  upload.zip
  inflating: hello.txt
*/
