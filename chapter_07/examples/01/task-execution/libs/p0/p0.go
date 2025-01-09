// libs/p0/p0.go

package p0

import (
    "fmt"
    "math/rand"
    "time"
)

// Task1 simulates task 1 execution
func Task1() {
    fmt.Println("Task1: Starting...")
    time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second) // Simulate random task duration
    fmt.Println("Task1: Finished.")
}

// Task2 simulates task 2 execution
func Task2() {
    fmt.Println("Task2: Starting...")
    time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second) // Simulate random task duration
    fmt.Println("Task2: Finished.")
}

// Task3 simulates task 3 execution
func Task3() {
    fmt.Println("Task3: Starting...")
    time.Sleep(time.Duration(rand.Intn(4)+3) * time.Second) // Simulate random task duration
    fmt.Println("Task3: Finished.")
}

// Task4 simulates task 4 execution with random rounds (1 to 500)
func Task4() {
    rounds := rand.Intn(500) + 1 // Random rounds between 1 and 500
    fmt.Printf("Task4: Starting with %d rounds...\n", rounds)

    for i := 1; i <= rounds; i++ {
        // Simulate work in each round
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+50)) // Random delay between 50ms to 150ms
        fmt.Printf("Task4: Round %d of %d\n", i, rounds)
    }

    fmt.Println("Task4: Finished.")
}

// ListTasks returns a list of all available tasks as strings
func ListTasks() []string {
    return []string{"Task1", "Task2", "Task3", "Task4"}
}
