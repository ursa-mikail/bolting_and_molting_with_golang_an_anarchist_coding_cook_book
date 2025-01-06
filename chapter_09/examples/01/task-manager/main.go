package main

import (
    "fmt"
    "task-manager/tasks"
    "task-manager/utils"
)

func main() {
    for {
        fmt.Println("\nTask Management System")
        fmt.Println("1. Add Task")
        fmt.Println("2. View Tasks")
        fmt.Println("3. Mark Task as Completed")
        fmt.Println("4. Export to CSV")
        fmt.Println("5. Exit")
        choice := utils.GetIntInput("Choose an option: ")

        switch choice {
        case 1:
            tasks.AddTask()
        case 2:
            tasks.ViewTasks()
        case 3:
            tasks.MarkTaskCompleted()
        case 4:
            tasks.ExportToCSV()
        case 5:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid option. Please try again.")
        }
    }
}

/*
% go mod init test-app
% go run main.go
*/
