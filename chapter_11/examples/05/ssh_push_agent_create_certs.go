package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
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
			return nil, fmt.Errorf("read key: %w", err)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("parse key: %w", err)
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

func uploadFile(sshClient *ssh.Client, localPath, remotePath string) error {
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	srcFile, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	info, _ := srcFile.Stat()
	writer, _ := session.StdinPipe()

	go func() {
		defer writer.Close()
		fmt.Fprintf(writer, "C0755 %d %s\n", info.Size(), filepath.Base(remotePath))
		io.Copy(writer, srcFile)
		fmt.Fprint(writer, "\x00")
	}()

	return session.Run("scp -t " + filepath.Dir(remotePath))
}

func runRemoteCommand(sshClient *ssh.Client, command string) (string, error) {
	session, err := sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var out bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out

	err = session.Run(command)
	return out.String(), err
}

func main() {
	host := os.Getenv("SSH_HOST") // e.g. "ubuntu@10.242.41.207"
	pass := os.Getenv("SSH_PASS") // optional
	key := os.Getenv("SSH_KEY")   // optional
	remoteDir := "~/cert_gen"
	localDownload := "certs_from_remote.zip"

	if host == "" {
		log.Fatal("Set SSH_HOST env var (e.g., ubuntu@host)")
	}
	parts := strings.SplitN(host, "@", 2)
	user, addr := parts[0], parts[1]

	cfg, err := getSSHConfig(user, pass, key)
	if err != nil {
		log.Fatal("SSH config error:", err)
	}

	client, err := ssh.Dial("tcp", addr+":22", cfg)
	if err != nil {
		log.Fatal("SSH dial failed:", err)
	}
	defer client.Close()

	// Step 1: Create remote folder
	fmt.Println("Creating remote directory...")
	if _, err := runRemoteCommand(client, "mkdir -p "+remoteDir); err != nil {
		log.Fatalf("mkdir failed: %v", err)
	}

	// Step 2: Upload Python script
	script := "gen_certs.py"
	fmt.Println("Uploading script...")
	remoteScript := remoteDir + "/" + script
	if err = uploadFile(client, script, remoteScript); err != nil {
		log.Fatalf("upload failed: %v", err)
	}

	// Step 3: Install Python deps
	setup := "sudo apt update -y && sudo apt install -y python3-pip zip && pip3 install cryptography"
	fmt.Println("Installing dependencies...")
	if _, err := runRemoteCommand(client, setup); err != nil {
		log.Println("Python3 setup error:", err)
	}

	// Step 4: Run Python script
	fmt.Println("Running Python cert generator on remote...")
	cmd := fmt.Sprintf("cd %s && python3 %s 10", remoteDir, script)
	output, err := runRemoteCommand(client, cmd)
	if err != nil {
		log.Fatal("Remote command failed:", err)
	}
	fmt.Println("Remote Output:\n" + output)

	// Step 5: Zip the generated certs
	fmt.Println("Zipping remote certs...")
	zipCmd := fmt.Sprintf("cd %s && zip -r certs.zip gen_certs", remoteDir)
	if _, err := runRemoteCommand(client, zipCmd); err != nil {
		log.Fatalf("zip failed: %v", err)
	}

	output, err = runRemoteCommand(client, "ls -lh "+remoteDir)
	fmt.Println("Files in remoteDir:\n", output)

	// Step 6: SHA256 remote
	fmt.Println("Calculating SHA256 on remote...")
	remoteShaCmd := fmt.Sprintf("openssl dgst -sha256 %s/certs.zip", remoteDir)

	remoteHash, err := runRemoteCommand(client, remoteShaCmd)
	if err != nil {
		log.Fatalf("Remote SHA256 failed: %v", err)
	}
	//fmt.Println("Remote SHA256: ", strings.Fields(remoteHash)[0])

	fields := strings.Fields(remoteHash)
	if len(fields) == 0 {
		log.Fatalf("SHA256 output was empty: %q", remoteHash)
	}
	fmt.Println("Remote SHA256: ", fields[0])

	// Step 7: Download zip
	fmt.Println("Downloading zip file...")
	remoteZip := fmt.Sprintf("%s/certs.zip", remoteDir)
	if err := scpDownloadFile(client, remoteZip, localDownload); err != nil {
		log.Fatalf("Download failed: %v", err)
	}

	// Step 8: SHA256 local
	localHash, err := computeSHA256(localDownload)
	if err != nil {
		log.Fatalf("Local SHA256 failed: %v", err)
	}
	fmt.Println("Local SHA256:  ", localHash)
}

func scpDownloadFile(sshClient *ssh.Client, remotePath, localPath string) error {
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var stdout bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stdout

	// Use cat to stream the file
	cmd := fmt.Sprintf("cat %s", remotePath)
	if err := session.Start(cmd); err != nil {
		return err
	}

	outFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, &stdout)
	return err
}

func computeSHA256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// ssh_push_agent_create_certs.go
// on the remote: sudo apt update -y && sudo apt install -y python3-pip zip openssl && pip3 install cryptography
// push_remote_agent.go
// openssl dgst -sha256 certs_from_remote.zip
