package tasks

// Task represents a task assigned to an individual
type Task struct {
	ID       int
	Assignee string
	Details  string
}

var tasks = make([]Task, 0)
var nextTaskID = 1
const maxTasks = 10

// AllTasksAssigned checks if the maximum tasks have been assigned
func AllTasksAssigned() bool {
	return len(tasks) >= maxTasks
}
