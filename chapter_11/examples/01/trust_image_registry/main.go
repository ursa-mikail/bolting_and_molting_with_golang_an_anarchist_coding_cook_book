package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
)

type SignedImage struct {
	ImageID   string `json:"image_id"`
	Issuer    string `json:"issuer"`
	Signature string `json:"signature"`
	Status    string `json:"status"`
	CVEURL    string `json:"cve_url,omitempty"`
}

var imageRegistry = struct {
	sync.Mutex
	Data map[string]SignedImage
}{Data: make(map[string]SignedImage)}

func generateSignature(imagePath string) (string, error) {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

func signImage(w http.ResponseWriter, r *http.Request) {
	var req SignedImage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	signature, err := generateSignature(req.ImageID) // Using ImageID as path
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusInternalServerError)
		return
	}

	req.Signature = signature
	req.Status = "Valid"

	imageRegistry.Lock()
	imageRegistry.Data[req.ImageID] = req
	imageRegistry.Unlock()

	json.NewEncoder(w).Encode(req)
}

func revokeImage(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ImageID string `json:"image_id"`
		CVEURL  string `json:"cve_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	imageRegistry.Lock()
	image, exists := imageRegistry.Data[req.ImageID]
	if !exists {
		imageRegistry.Unlock()
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	image.Status = "Revoked"
	image.CVEURL = req.CVEURL
	imageRegistry.Data[req.ImageID] = image
	imageRegistry.Unlock()

	notifyIssuer(image)
	json.NewEncoder(w).Encode(image)
}

func notifyIssuer(image SignedImage) {
	fmt.Printf("Notifying issuer %s: %s\n", image.Issuer, toJSON(image))
}

func toJSON(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "Unknown"
}

func main() {
	http.HandleFunc("/sign", signImage)
	http.HandleFunc("/revoke", revokeImage)

	localIP := getLocalIP()
	port := "8080"
	fmt.Printf("Server running on http://%s:%s\n", localIP, port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
