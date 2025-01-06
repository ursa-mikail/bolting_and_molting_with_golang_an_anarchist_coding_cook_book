package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetStringInput reads a string input from the user
func GetStringInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// GetIntInput reads an integer input from the user
func GetIntInput(prompt string) int {
	for {
		fmt.Print(prompt)
		input := GetStringInput("")
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
}

