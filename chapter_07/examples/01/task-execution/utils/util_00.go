// utils/util_00.go

package utils

import (
    "fmt"
    "time"
)

// LogExecutionTime logs the execution time of a task
func LogExecutionTime(taskName string, start time.Time) {
    end := time.Now()
    duration := end.Sub(start)
    fmt.Printf("Task %s: Started at %v, Ended at %v, Duration: %v\n", taskName, start, end, duration)
}
