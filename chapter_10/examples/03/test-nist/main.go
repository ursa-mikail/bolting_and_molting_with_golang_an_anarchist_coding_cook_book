package main

import (
    "log"
    "test-nist/libs"
)

func main() {
    url := "https://raw.githubusercontent.com/usnistgov/ACVP-Server/master/gen-val/json-files/ACVP-AES-GCM-1.0/internalProjection.json"
    
    test, err := libs.FetchNISTTest(url)
    if err != nil {
        log.Fatalf("Error fetching NIST test vectors: %v", err)
    }

    libs.RunTest(test)
}



/*
% go mod init nist-test
% go run main.go

Error Handling: The errors are logged with details to identify failures.
Comparison: Ciphertext and tag are compared with the expected values.
Logging: Detailed logs ensure traceability of test results.


Running tests for Algorithm: AES
Mode: GCM

Test Case 1:
Test case 1 passed!
Decrypted Plaintext: Hello, NIST!

Test Case 2:
Test case 2 failed: Mismatch in ciphertext or tag

*/
