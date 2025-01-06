package tasks

import (
	"fmt"
)

// AssignTask assigns a new task to the task list
func AssignTask(assignee string, details string) {
	task := Task{
		ID:       nextTaskID,
		Assignee: assignee,
		Details:  details,
	}
	tasks = append(tasks, task)
	nextTaskID++

	fmt.Printf("Task assigned to %s: %s (Task ID: %d)\n", assignee, details, task.ID)
	fmt.Printf("%d tasks remaining to assign.\n", maxTasks-len(tasks))
}

// GreetUsers welcomes the user to the task assignment system
func GreetUsers() {
	fmt.Printf("Welcome to the Task Assignment System.\n")
	fmt.Printf("You can assign up to %d tasks.\n", maxTasks)
}

// ValidateUserInput validates the assignee name and task details
func ValidateUserInput(assignee string, details string) (bool, bool) {
	isValidAssignee := len(assignee) >= 2
	isValidTask := len(details) >= 5
	return isValidAssignee, isValidTask
}
