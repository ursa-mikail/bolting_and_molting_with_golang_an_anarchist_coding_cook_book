// main.go

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
    "task-execution/libs/p0"
    "task-execution/utils"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // List available tasks
    tasks := p0.ListTasks()
    fmt.Println("Available tasks:", tasks)

    // Prepare for concurrent execution
    var wg sync.WaitGroup

    // Execute random tasks concurrently
    for i := 0; i < 5; i++ {
        wg.Add(1)

        // Randomly select a task function to execute
        go func(taskName string) {
            defer wg.Done()

            // Log the task start time
            startTime := time.Now()

            // Call the task function based on the task name
            switch taskName {
            case "Task1":
                p0.Task1()
            case "Task2":
                p0.Task2()
            case "Task3":
                p0.Task3()
            case "Task4":
                p0.Task4()
            default:
                fmt.Println("Unknown task:", taskName)
            }

            // Log the task execution time
            utils.LogExecutionTime(taskName, startTime)
        }(tasks[rand.Intn(len(tasks))]) // Randomly choose a task
    }

    // Wait for all goroutines to finish
    wg.Wait()

    fmt.Println("All tasks completed.")
}
