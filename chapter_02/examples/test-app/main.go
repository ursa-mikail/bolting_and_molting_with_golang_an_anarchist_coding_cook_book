package main

import (
    "fmt"
    "test-app/libs/p0"
    "test-app/utils"
)

func main() {
    fmt.Print(utils.Hello() + "\n")

    // Introduce the theme
    fmt.Println("Welcome to 'Commitment Issues: Variables and Data Types'!")
    fmt.Println("Today, we'll explore Golang's quirky variable and data type relationships.")

    // Call ExplainVariables
    p0.ExplainVariables()

    // Variables with different data types
    var name string = "Gopher"
    age := 10
    height := 1.75
    isGopherCute := true

    // Print them out
    fmt.Println("\nMeet our star:")
    fmt.Printf("- Name: %s (string)\n", name)
    fmt.Printf("- Age: %d (int)\n", age)
    fmt.Printf("- Height: %.2f (float64)\n", height)
    fmt.Printf("- Is Gopher cute? %t (bool)\n", isGopherCute)

    // Show interaction between types
    fmt.Println("\nBut wait, what happens if we try mixing them?")
    fmt.Println("- Gopher says, 'Hey, can I combine age and height?'")
    result := float64(age) + height
    fmt.Printf("Result: %.2f (age + height)\n", result)

    // Generate SHA256 IDs
    fmt.Println("\nGenerating SHA256 IDs for some favorite variables:")
    inputs := []string{name, fmt.Sprintf("%d", age), fmt.Sprintf("%.2f", height), fmt.Sprintf("%t", isGopherCute)}
    ids := p0.GenerateSHA256IDs(inputs)

    for i, id := range ids {
        fmt.Printf("- %s: %s\n", inputs[i], id)
    }

    // Play with constants
    const motto string = "Keep coding and stay quirky!"
    fmt.Printf("\nA constant reminder: %s\n", motto)

    // End with a laugh
    fmt.Println("\nAnd remember: Variables may have commitment issues, but constants are forever.")
}


/*
% go mod init test-app
% go run main.go
*/
