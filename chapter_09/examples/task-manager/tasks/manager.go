package tasks

import (
	"fmt"
	"task-manager/utils"
)

// AddTask adds a new task to the list
func AddTask() {
	description := utils.GetStringInput("Enter task description: ")

	var priority string
	for {
		priority = utils.GetStringInput("Enter task priority (High, Medium, Low): ")
		priority = NormalizePriority(priority)
		if priority != "" {
			break
		}
		fmt.Println("Invalid priority. Please choose High, Medium, or Low.")
	}

	task := Task{
		ID:          nextID,
		Description: description,
		Status:      "Pending",
		Priority:    priority,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully.")
}

// ViewTasks displays the current list of tasks
func ViewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\nCurrent Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d | Description: %s | Status: %s | Priority: %s\n",
			task.ID, task.Description, task.Status, task.Priority)
	}
}

// MarkTaskCompleted marks a task as completed
func MarkTaskCompleted() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available to mark as completed.")
		return
	}

	id := utils.GetIntInput("Enter the task ID to mark as completed: ")

	for i, task := range tasks {
		if task.ID == id {
			if task.Status == "Completed" {
				fmt.Println("Task is already completed.")
				return
			}
			tasks[i].Status = "Completed"
			fmt.Println("Task marked as completed.")
			return
		}
	}
	fmt.Println("Task ID not found.")
}

