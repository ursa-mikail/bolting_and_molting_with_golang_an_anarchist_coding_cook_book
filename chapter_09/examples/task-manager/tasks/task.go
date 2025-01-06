package tasks

import "strings"

// Task represents a single task
type Task struct {
	ID          int
	Description string
	Status      string // "Pending" or "Completed"
	Priority    string // "High", "Medium", "Low"
}

var tasks = make([]Task, 0)
var nextID = 1

// NormalizePriority ensures priority is valid
func NormalizePriority(priority string) string {
	priority = strings.Title(strings.ToLower(priority))
	if priority != "High" && priority != "Medium" && priority != "Low" {
		return ""
	}
	return priority
}
