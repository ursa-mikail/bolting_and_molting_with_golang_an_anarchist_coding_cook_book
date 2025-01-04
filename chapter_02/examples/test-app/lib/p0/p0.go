package p0

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

// GenerateSHA256IDs demonstrates SHA256 ID generation for strings.
func GenerateSHA256IDs(inputs []string) []string {
    ids := []string{}
    for _, input := range inputs {
        hash := sha256.Sum256([]byte(input))
        ids = append(ids, hex.EncodeToString(hash[:]))
    }
    return ids
}

// ExplainVariables provides a witty description of variable types.
func ExplainVariables() {
    fmt.Println("Commitment Issues: Why can't variables and types just get along?")
    fmt.Println("- A string says, 'I'm here for the long haul, but only with text.'")
    fmt.Println("- An int replies, 'Numbers are my thing, and I prefer no decimal drama.'")
    fmt.Println("- A float counters, 'Decimals? Drama? No, it's elegance.'")
    fmt.Println("- A bool smirks, 'True or false, I'm all about commitment.'")
}
