package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput collects assignee name and task details from the user
func GetUserInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the assignee name:")
	assignee, _ := reader.ReadString('\n')
	assignee = strings.TrimSpace(assignee)

	fmt.Println("Enter the task details:")
	details, _ := reader.ReadString('\n')
	details = strings.TrimSpace(details)

	return assignee, details
}
