
We automate the scaffolding of the <a href="https://github.com/ursa-mikail/golang-gaia-basic-structure/tree/main"> golang-gaia-basic-structure</a>.

<pre>
chmod +x make_go.sh
# Run the script with your desired module name:
# ./make_go.sh example.com/demo
./make_go.sh test-app

For [01]:
# Resulting Structure
After running the script, the structure will look like this:

task-manager 
├── go.mod
├── libs
│   └── p0
│       └── p0.go
├── main.go
└── utils
    └── util_00.go

# modify the generated main.go
# create dir: tasks/
# and add the *.go
The structure will look like this:

task-manager/
├── main.go
├── tasks/
│   ├── manager.go
│   ├── csv_export.go
│   └── task.go
├── utils/
│   └── input.go
└── tasks.csv


# test run:
% cd task-manager
% go run main.go

out:
</pre>
```
Task Management System
1. Add Task
2. View Tasks
3. Mark Task as Completed
4. Export to CSV
5. Exit
Choose an option: 1

Enter task description: Clean the kitchen
Enter task priority (High, Medium, Low): High
Task added successfully.

```

CSV Output (tasks.csv):
```
ID,Description,Status,Priority
1,Clean the kitchen,Pending,High
```

<pre>
chmod +x make_go.sh
# Run the script with your desired module name:
# ./make_go.sh example.com/demo
./make_go.sh test-app

For [02]:
# Resulting Structure
task-assignment/
├── main.go
├── tasks/
│   ├── task.go
│   ├── manager.go
│   ├── input.go

out:
</pre>
```
Welcome to the Task Assignment System.
You can assign up to 10 tasks.

Enter the assignee name:
Alice
Enter the task details:
Complete the quarterly report.
Task assigned to Alice: Complete the quarterly report (Task ID: 1)
9 tasks remaining to assign.

Enter the assignee name:
Bob
Enter the task details:
Prepare the presentation for the client meeting.
Task assigned to Bob: Prepare the presentation for the client meeting (Task ID: 2)
8 tasks remaining to assign.

```


### Advantages of This Structure
Modularity: Different aspects of the program are separated into logical packages (tasks and utils).
Reusability: Code can be reused across other programs or expanded easily.
Scalability: Adding new features, such as task categories or deadlines, becomes straightforward.
Ease of Maintenance: Each package has a clear responsibility, making debugging and updates more efficient.

