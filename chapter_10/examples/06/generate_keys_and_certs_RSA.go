package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	// Generate RSA Private Key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Create Certificate Template
	template := x509.Certificate{
		SerialNumber: big.NewInt(20250516001),
		Subject: pkix.Name{
			Organization:  []string{"Your Org Name"},
			Country:       []string{"US"},
			Province:      []string{"California"},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Market Street"},
			PostalCode:    []string{"94103"},
			CommonName:    "your-domain.com",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // valid for 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Create self-signed certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	// Encode and write certificate to file
	certFile, err := os.Create("certificate.crt")
	if err != nil {
		panic(err)
	}
	defer certFile.Close()
	pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	// Encode and write private key to file
	keyFile, err := os.Create("decryption-key-server.key")
	if err != nil {
		panic(err)
	}
	defer keyFile.Close()
	privateKeyDER := x509.MarshalPKCS1PrivateKey(privateKey)
	pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyDER})

	// Parse the certificate for debug output
	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		panic(err)
	}

	// --- Debug Output Section ---
	fmt.Println("\n🔍 --- CERTIFICATE DETAILS ---")
	fmt.Printf("Subject: %s\n", cert.Subject.String())
	fmt.Printf("Issuer: %s\n", cert.Issuer.String())
	fmt.Printf("Serial Number: %s\n", cert.SerialNumber.String())
	fmt.Printf("Not Before: %s\n", cert.NotBefore)
	fmt.Printf("Not After: %s\n", cert.NotAfter)
	fmt.Printf("Is CA: %v\n", cert.IsCA)
	fmt.Printf("Key Usage: %v\n", cert.KeyUsage)
	fmt.Printf("Ext Key Usage: %v\n", cert.ExtKeyUsage)
	fmt.Printf("Signature Algorithm: %v\n", cert.SignatureAlgorithm)
	fmt.Printf("Public Key Algorithm: %v\n", cert.PublicKeyAlgorithm)

	// Print the certificate in PEM format
	fmt.Println("\n📄 --- PEM ENCODED CERTIFICATE ---")
	pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	// --- Print Private Key in HEX and BASE64 ---
	fmt.Println("\n🔐 --- PRIVATE KEY RAW (HEX) ---")
	fmt.Println(hex.EncodeToString(privateKeyDER))

	fmt.Println("\n🔐 --- PRIVATE KEY RAW (BASE64) ---")
	fmt.Println(base64.StdEncoding.EncodeToString(privateKeyDER))
}

/*
% go mod tidy
% go run generate_keys_and_certs_RSA.go


🔍 --- CERTIFICATE DETAILS ---
Subject: CN=your-domain.com,O=Your Org Name,POSTALCODE=94103,STREET=Market Street,L=San Francisco,ST=California,C=US
Issuer: CN=your-domain.com,O=Your Org Name,POSTALCODE=94103,STREET=Market Street,L=San Francisco,ST=California,C=US
Serial Number: 20250516001
Not Before: 2009-11-10 23:00:00 +0000 UTC
Not After: 2010-11-10 23:00:00 +0000 UTC
Is CA: true
Key Usage: 5
Ext Key Usage: [1]
Signature Algorithm: SHA256-RSA
Public Key Algorithm: RSA

📄 --- PEM ENCODED CERTIFICATE ---
-----BEGIN CERTIFICATE-----
MIIEADCCAuigAwIBAgIFBLcGWiEwDQYJKoZIhvcNAQELBQAwgZQxCzAJBgNVBAYT
AlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2Nv
MRYwFAYDVQQJEw1NYXJrZXQgU3RyZWV0MQ4wDAYDVQQREwU5NDEwMzEWMBQGA1UE
ChMNWW91ciBPcmcgTmFtZTEYMBYGA1UEAxMPeW91ci1kb21haW4uY29tMB4XDTA5
MTExMDIzMDAwMFoXDTEwMTExMDIzMDAwMFowgZQxCzAJBgNVBAYTAlVTMRMwEQYD
VQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRYwFAYDVQQJ
Ew1NYXJrZXQgU3RyZWV0MQ4wDAYDVQQREwU5NDEwMzEWMBQGA1UEChMNWW91ciBP
cmcgTmFtZTEYMBYGA1UEAxMPeW91ci1kb21haW4uY29tMIIBIjANBgkqhkiG9w0B
AQEFAAOCAQ8AMIIBCgKCAQEAuUMw2EJLEqCqHYOaR/tmG0wfE1b3yDAu/taMyMg5
agyYk8FzwdAM9RbvIHfdBRQolq7DudvFGt1g1IT/dLXjRxZsi0ZmbGBGi8WYRNV7
vYoDXoTNn7F+h1cJ22JJB6kKWe0WPXIdnHNUGvN2io7Uq0MPNYZwh9IfAZl+HRF9
kPvnq1psWmLGGYOqKdeulXvjPBlf6d5xMZFbiKbLgvfRE5l+cp89OrHZux36FDZS
CmY6tLGq7VkL09Tq8cbtDRRQEMkaAVb3FSz31XKo+Rfvjkg5qxAuA1r8EedzNlgv
Aak72vxgdR0DTBNR4iQ0QHBMWUoPerEQc6g4czXeInrhgQIDAQABo1cwVTAOBgNV
HQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB
/zAdBgNVHQ4EFgQUZQ22hy2oAH7kiW7UbkyjG2c1gncwDQYJKoZIhvcNAQELBQAD
ggEBABQRBjbkkz0xJ/zHrydybYWFzGRJbFSUr8oCURtxNdNNwKcVFC4SwvMq1mrJ
8NKgpkUHTr2RhHGfjgjz20GPAZ616JQ6LP7Hvs/ySy2Kpq9Tyu8QR0OquEYK8OYU
rrJvtqVXeCy5piSRDBI08Wdi+brbbMDnwjS5s2kgfwk9+a0vQJhy4TME50LhCLuz
OrKpZbOOPWj7JrX+jy9pjNxxcZ3hSqNmdsff3kpBKJqANLq9Tv+povwddsmkT1BF
QjSxyRSHlvcLpnpUVVwuorRgBl+EE6NRHl2K06le57WwkzmQDBMzAUe2G3H74fGd
eSdon32ozQVSuB3XU7Ek+hvFWo0=
-----END CERTIFICATE-----

🔐 --- PRIVATE KEY RAW (HEX) ---
308204a50201000282010100b94330d8424b12a0aa1d839a47fb661b4c1f1356f7c8302efed68cc8c8396a0c9893c173c1d00cf516ef2077dd05142896aec3b9dbc51add60d484ff74b5e347166c8b46666c60468bc59844d57bbd8a035e84cd9fb17e875709db624907a90a59ed163d721d9c73541af3768a8ed4ab430f35867087d21f01997e1d117d90fbe7ab5a6c5a62c61983aa29d7ae957be33c195fe9de7131915b88a6cb82f7d113997e729f3d3ab1d9bb1dfa1436520a663ab4b1aaed590bd3d4eaf1c6ed0d145010c91a0156f7152cf7d572a8f917ef8e4839ab102e035afc11e77336582f01a93bdafc60751d034c1351e2243440704c594a0f7ab11073a8387335de227ae1810203010001028201001748dfa1beab4adb10a686fce4212fd32c7d17642dde6657f1c5ab8a2cc8329104b8456f3423df0d0ae5e131cf80ace8f10f64203f0405fd5f0736f4188a27d7ae74d408742518f0a4ee9fd5bef33f7e52301606f6144befdcbbab6c2dfb9d7e64910ce38a2d3895b23374f828300d80bcfbf36143517ed250d426eaa6611ed0a3a325e8c522ee913185c8b82d0db5f01c6181f6bd366e628bba6a975fb93dd9202288c6e323ea18f939ae4d95c175ab77fb4505d3bd7935ab3a3be9348b4e552d28ff5d95999badb1f53eedd22d5349cc378a2472538f9d8e66c1b7041cceeb5a87785a4c3c92d5121f603ca2eba2933a57939843c91fc342956a1571ffc8c502818100e2954daf9207b1afe0eff77a625e76f674910e1fc14ab4021c507f7e2d31ded535dbc4f5662b5c7433084e1c033f8aa1af025132edca381921d61064b092b772dfb3473a84ba3a3a850f275c1d1b97e464d2dde209d11d5dd64a95e6ae0fa894228efe783b201b0baf961ede19b98280f9c840945ff034fe2a216e99f77e04ab02818100d1508e17885f70a6ebd394bb49db3e0dde7324cf64510cedd152feffb0718465ec2d634f076ce7e865791be8a87511bfd51d3a2d12d15699863cb84090308da4bc88558a47a63491fd53ed08e2b9b1aa56b9ed68edc6b54cfdb96443b452690563cca1f36590a7edb064634502ce1db911e87efff20ad85ed5548610fa4f7a8302818100a40c06ffdb917883672f3615448d324243a915f9c5bf694eb1e3523b2f4b9112faf88fc440619a8182c53b902a85e1c027ee5c8dabfe21d98856eced89addfe5eab2691aa42814aa73d5c1e07e2e912d4a162591571c30cfd4ab91963f6594df139a46011485ad2f2bc45bbae01320ba729e1aab923e5e8fba3f144da313791b02818100bfda9c4309a6a134fe97528e5065fa7cefd68e78e1b440ab82e01606fb0d3193d264cfcfd2388a6043123cf885dfb84322edc533273a65ac169475360873d88094a649f19336ce9cb99c417991872f47b872771f64426c4f838878546e65d30933ae8f0aef9f25ad01af22265129e7a888b79820b51427bad4c2c297b137444d028181009f0b8603b16fe5674d6c741fa4eec1785dab3ebae4ee7f75aff02389cffb050f6a0c8db5b75846724826a40b052f2048228e3bad4e2bcc73a700157b9db5e54267c18e33f98ab2c594fa1f6b3c51594f31a915dfab1c78ea5a9d9e48fe5ddf683ffee21dc1d49d0998c6c7f64d67b7814ec822374f6971f2a0c3faf547726867

🔐 --- PRIVATE KEY RAW (BASE64) ---
MIIEpQIBAAKCAQEAuUMw2EJLEqCqHYOaR/tmG0wfE1b3yDAu/taMyMg5agyYk8FzwdAM9RbvIHfdBRQolq7DudvFGt1g1IT/dLXjRxZsi0ZmbGBGi8WYRNV7vYoDXoTNn7F+h1cJ22JJB6kKWe0WPXIdnHNUGvN2io7Uq0MPNYZwh9IfAZl+HRF9kPvnq1psWmLGGYOqKdeulXvjPBlf6d5xMZFbiKbLgvfRE5l+cp89OrHZux36FDZSCmY6tLGq7VkL09Tq8cbtDRRQEMkaAVb3FSz31XKo+Rfvjkg5qxAuA1r8EedzNlgvAak72vxgdR0DTBNR4iQ0QHBMWUoPerEQc6g4czXeInrhgQIDAQABAoIBABdI36G+q0rbEKaG/OQhL9MsfRdkLd5mV/HFq4osyDKRBLhFbzQj3w0K5eExz4Cs6PEPZCA/BAX9Xwc29BiKJ9eudNQIdCUY8KTun9W+8z9+UjAWBvYUS+/cu6tsLfudfmSRDOOKLTiVsjN0+CgwDYC8+/NhQ1F+0lDUJuqmYR7Qo6Ml6MUi7pExhci4LQ218Bxhgfa9Nm5ii7pql1+5PdkgIojG4yPqGPk5rk2VwXWrd/tFBdO9eTWrOjvpNItOVS0o/12VmZutsfU+7dItU0nMN4okclOPnY5mwbcEHM7rWod4Wkw8ktUSH2A8ouuikzpXk5hDyR/DQpVqFXH/yMUCgYEA4pVNr5IHsa/g7/d6Yl529nSRDh/BSrQCHFB/fi0x3tU128T1ZitcdDMIThwDP4qhrwJRMu3KOBkh1hBksJK3ct+zRzqEujo6hQ8nXB0bl+Rk0t3iCdEdXdZKleauD6iUIo7+eDsgGwuvlh7eGbmCgPnIQJRf8DT+KiFumfd+BKsCgYEA0VCOF4hfcKbr05S7Sds+Dd5zJM9kUQzt0VL+/7BxhGXsLWNPB2zn6GV5G+iodRG/1R06LRLRVpmGPLhAkDCNpLyIVYpHpjSR/VPtCOK5sapWue1o7ca1TP25ZEO0UmkFY8yh82WQp+2wZGNFAs4duRHofv/yCthe1VSGEPpPeoMCgYEApAwG/9uReINnLzYVRI0yQkOpFfnFv2lOseNSOy9LkRL6+I/EQGGagYLFO5AqheHAJ+5cjav+IdmIVuztia3f5eqyaRqkKBSqc9XB4H4ukS1KFiWRVxwwz9SrkZY/ZZTfE5pGARSFrS8rxFu64BMgunKeGquSPl6Puj8UTaMTeRsCgYEAv9qcQwmmoTT+l1KOUGX6fO/WjnjhtECrguAWBvsNMZPSZM/P0jiKYEMSPPiF37hDIu3FMyc6ZawWlHU2CHPYgJSmSfGTNs6cuZxBeZGHL0e4cncfZEJsT4OIeFRuZdMJM66PCu+fJa0BryImUSnnqIi3mCC1FCe61MLCl7E3RE0CgYEAnwuGA7Fv5WdNbHQfpO7BeF2rPrrk7n91r/Ajic/7BQ9qDI21t1hGckgmpAsFLyBIIo47rU4rzHOnABV7nbXlQmfBjjP5irLFlPofazxRWU8xqRXfqxx46lqdnkj+Xd9oP/7iHcHUnQmYxsf2TWe3gU7IIjdPaXHyoMP69UdyaGc=

*/