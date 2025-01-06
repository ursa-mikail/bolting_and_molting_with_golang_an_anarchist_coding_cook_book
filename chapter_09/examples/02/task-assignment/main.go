package main

import (
    "fmt"
    "task-assignment/tasks"
)

func main() {
    tasks.GreetUsers()

    for {
        assignee, taskDetails := tasks.GetUserInput()
        isValidAssignee, isValidTask := tasks.ValidateUserInput(assignee, taskDetails)

        if isValidAssignee && isValidTask {
            tasks.AssignTask(assignee, taskDetails)
            fmt.Println("Task assignment successful.")

            if tasks.AllTasksAssigned() {
                fmt.Println("All tasks have been assigned.")
                break
            }
        } else {
            if !isValidAssignee {
                fmt.Println("The assignee name is too short or invalid.")
            }
            if !isValidTask {
                fmt.Println("The task description is too short or invalid.")
            }
            continue
        }
    }
}
