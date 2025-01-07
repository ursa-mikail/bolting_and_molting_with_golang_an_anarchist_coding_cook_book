package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "task-assignment/tasks"
)

func main() {
    tasks.GreetUsers()
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Choose an option:")
        fmt.Println("1. Assign a Task")
        fmt.Println("2. List All Tasks")
        fmt.Println("3. Exit")
        fmt.Print("Enter your choice: ")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            assignee, taskDetails := tasks.GetUserInput()
            isValidAssignee, isValidTask := tasks.ValidateUserInput(assignee, taskDetails)

            if isValidAssignee && isValidTask {
                tasks.AssignTask(assignee, taskDetails)
            } else {
                if !isValidAssignee {
                    fmt.Println("The assignee name is too short or invalid.")
                }
                if !isValidTask {
                    fmt.Println("The task description is too short or invalid.")
                }
            }

        case "2":
            tasks.ListAllTasks()

        case "3":
            fmt.Println("Exiting the Task Assignment System. Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }

        if tasks.AllTasksAssigned() {
            fmt.Println("All tasks have been assigned. Exiting.")
            break
        }
    }
}
