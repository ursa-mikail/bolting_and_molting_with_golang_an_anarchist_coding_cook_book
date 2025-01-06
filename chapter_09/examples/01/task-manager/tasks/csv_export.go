package tasks

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ExportToCSV exports the task list to a CSV file
func ExportToCSV() {
	file, err := os.Create("tasks.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Description", "Status", "Priority"})
	for _, task := range tasks {
		writer.Write([]string{
			strconv.Itoa(task.ID),	// strconv.FormatUint(uint64(task.ID), 10) // base-10
			task.Description,
			task.Status,
			task.Priority,
		})
	}
	fmt.Println("Tasks exported to tasks.csv successfully.")
}
