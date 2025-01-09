package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())

    // Create a channel for communication between threads
    messageChannel := make(chan string)

    // Start the first thread
    go func() {
        for i := 0; i < 5; i++ {
            message := fmt.Sprintf("Thread 1 says: Message %d", i+1)
            fmt.Println(message)
            messageChannel <- message // Send message to the other thread
            time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500)) // Random delay
        }
        close(messageChannel) // Close the channel when done
    }()

    // Start the second thread
    go func() {
        for message := range messageChannel {
            fmt.Printf("Thread 2 received: %s\n", message)
            response := fmt.Sprintf("Thread 2 responds to: %s", message)
            fmt.Println(response)
            time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500)) // Random delay
        }
    }()

    // Wait for threads to complete
    time.Sleep(5 * time.Second)
    fmt.Println("Conversation between threads completed.")
}


/*
% go mod init test-app
% go run main.go
*/
