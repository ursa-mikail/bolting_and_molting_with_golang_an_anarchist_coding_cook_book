package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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

func cloneAndZipRepo(repoURL, zipName string) (string, string, error) {
	// Clone repo
	repoName := strings.TrimSuffix(filepath.Base(repoURL), ".git")
	if err := exec.Command("git", "clone", repoURL).Run(); err != nil {
		return "", "", fmt.Errorf("git clone failed: %w", err)
	}
	defer os.RemoveAll(repoName)

	// Create zip
	zipFile, err := os.Create(zipName)
	if err != nil {
		return "", "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	err = filepath.WalkDir(repoName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		relPath, _ := filepath.Rel(filepath.Dir(repoName), path)
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		fw, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, file)
		return err
	})
	if err != nil {
		return "", "", err
	}
	zipWriter.Close()

	// Calculate SHA256 of the zip file
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

	logDir := filepath.Dir(remotePath)
	logFile := filepath.Join(logDir, "deploy_log.txt")
	timestamp := time.Now().Format(time.RFC3339)

	// Enhanced verification script with logging
	cmds := fmt.Sprintf(`
cd %s && 
actual=$(sha256sum %s | awk '{print $1}') && 
expected="%s" && 
if [ "$actual" = "$expected" ]; then 
    echo "[%s] SUCCESS: SHA256 match confirmed. Expected: $expected, Actual: $actual" >> %s && 
    unzip -o %s && 
    echo "[%s] Files extracted successfully" >> %s && 
    echo "Verification passed - files extracted"; 
else 
    echo "[%s] ERROR: SHA256 mismatch! Expected: $expected, Actual: $actual" >> %s && 
    echo "Verification failed - SHA256 mismatch"; 
    exit 1; 
fi
`, logDir, filepath.Base(remotePath), remoteHash, timestamp, logFile,
		filepath.Base(remotePath), timestamp, logFile, timestamp, logFile)

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

	// Log deployment result locally as well
	localLog := fmt.Sprintf("[%s] Deployment to %s - SHA256: %s - Result: %v\n",
		timestamp, remotePath, remoteHash, err == nil)

	fmt.Print(localLog)

	// Optionally write to a local log file
	f, logErr := os.OpenFile("local_deploy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if logErr == nil {
		defer f.Close()
		f.WriteString(localLog)
	}

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
	repoURL := os.Getenv("REPO_URL")
	remotePath := os.Getenv("REMOTE_PATH")

	if host == "" {
		log.Fatal("Missing SSH_HOST")
	}

	if repoURL == "" {
		repoURL = "https://github.com/ursa-mikail/mechanisms" // Default if not provided
	}

	if remotePath == "" {
		remotePath = "~/mechanisms.zip" // Default if not provided
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

	zipName := filepath.Base(repoURL) + ".zip"
	zipName, localHash, err := cloneAndZipRepo(repoURL, zipName)
	if err != nil {
		log.Fatal("Clone and zip failed:", err)
	}

	fmt.Println("Local SHA256:", localHash)

	err = uploadAndVerify(zipName, remotePath, localHash, client)
	if err != nil {
		log.Fatal("Upload or verification failed:", err)
	} else {
		fmt.Println("Deployment completed successfully")
	}
}

/* for SSH variables, refer: https://github.com/ursa-mikail/shell_script_utility/blob/main/scripts/utilities/dev_shell.sh
% go mod tidy
% go run main.go
Local SHA256: 21461a5a51467e4661765df42df09b5555a7f09d0bfd877a2ab39b415f28ca0f
Remote Output:
 Archive:  mechanisms.zip
  inflating: mechanisms/.git/HEAD
 :
  inflating: mechanisms/security/integrity/python/object_signing.py
Verification passed - files extracted

[2025-05-10T09:51:35-07:00] Deployment to ~/mechanisms.zip - SHA256: 21461a5a51467e4661765df42df09b5555a7f09d0bfd877a2ab39b415f28ca0f - Result: true
Deployment completed successfully
*/
