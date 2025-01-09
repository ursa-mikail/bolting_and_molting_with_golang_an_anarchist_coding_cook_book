
We automate the scaffolding of the <a href="https://github.com/ursa-mikail/golang-gaia-basic-structure/tree/main"> golang-gaia-basic-structure</a>.

<pre>
Each round involves a random sleep duration to simulate the work being done.
The task functions are randomly selected for concurrent execution, and their execution times are logged.
The program uses Go's concurrency model to execute multiple tasks in parallel.

Task Functions:
libs/p0/p0.go contains 4 task functions (Task1, Task2, and Task3, Task4) that simulate some work with a random duration. The ListTasks function returns a list of these tasks' names. 

Task4 runs a random number of rounds (between 1 and 500). In each round, it simulates some work by sleeping for a random period between 50 milliseconds and 150 milliseconds.

Concurrency:
The main.go file uses goroutines (go func()) to execute tasks concurrently. Each goroutine executes a randomly selected task from the list returned by p0.ListTasks().
Task Execution Time:

The utils/util_00.go file provides a utility function LogExecutionTime, which logs the start time, end time, and duration of each task.

Random Task Selection:
In the main.go, tasks are randomly selected using rand.Intn(len(tasks)), and the selected task is executed in a goroutine. 

Concurrency Synchronization:
The sync.WaitGroup is used to ensure that the main goroutine waits for all tasks to finish before exiting.

Concurrency and Logging:
Each selected task is executed concurrently using goroutines, and the execution time for each task is logged using the LogExecutionTime utility from utils/util_00.go.

</pre>

<pre>
chmod +x make_go.sh
# Run the script with your desired module name:
# ./make_go.sh example.com/demo
./make_go.sh test-app

# Resulting Structure
After running the script, the structure will look like this:

task-execution
├── go.mod
├── libs
│   └── p0
│       └── p0.go
├── main.go
└── utils
    └── util_00.go

# modify the generated main.go, p0.go, util_00.go

# test run:
% cd task-execution
% go run main.go

out:
</pre>
```
Available tasks: [Task1 Task2 Task3]
Task2: Starting...
Task3: Starting...
Task1: Starting...
Task3: Finished.
Task2: Finished.
Task1: Finished.
Task Task1: Started at 2025-01-08 14:01:00.100234 +0000 UTC, Ended at 2025-01-08 14:01:01.101234 +0000 UTC, Duration: 1.001000s
Task Task3: Started at 2025-01-08 14:01:00.100234 +0000 UTC, Ended at 2025-01-08 14:01:03.103234 +0000 UTC, Duration: 3.003000s
Task Task2: Started at 2025-01-08 14:01:00.100234 +0000 UTC, Ended at 2025-01-08 14:01:02.102234 +0000 UTC, Duration: 2.002000s
All tasks completed.

```