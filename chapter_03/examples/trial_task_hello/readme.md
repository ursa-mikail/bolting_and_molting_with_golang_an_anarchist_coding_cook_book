"""
# Trial Task Hello

A simple Go CLI program with task-based automation using `go-task`.

## 🛠️ Prerequisites

- [Go](https://golang.org/dl/)
- [Task](https://taskfile.dev/) (install via `brew install go-task/tap/go-task`)

✅ Verify installation: `task --version`


## 📁 Project Structure

```
.
├── main.go
├── Taskfile.yml
└── README.md
```

## 🚀 Usage

### Build the executable

```bash
task build
```

### Run in Hello mode (dev environment)

```bash
task hello
```

### Run in Goodbye mode (prod environment)

```bash
task goodbye
```

### Run Hello in production

```bash
task hello:prod
```

### Run Goodbye in development

```bash
task goodbye:dev
```

### Clean the binary

```bash
task clear_executable
```

## 📄 Taskfile Commands

```yaml
version: '3'

tasks:
  build:
    desc: Build the Go binary
    cmds:
      - go build -o trial_task_hello main.go

  hello:
    desc: Run Hello mode (default env = dev)
    deps: [build]
    cmds:
      - ./trial_task_hello -mode=hello -env=dev

  goodbye:
    desc: Run Goodbye mode (env = prod)
    deps: [build]
    cmds:
      - ./trial_task_hello -mode=goodbye -env=prod

  hello:prod:
    desc: Run Hello in production
    deps: [build]
    cmds:
      - ./trial_task_hello -mode=hello -env=prod

  goodbye:dev:
    desc: Run Goodbye in development
    deps: [build]
    cmds:
      - ./trial_task_hello -mode=goodbye -env=dev

  clear_executable:
    desc: Remove the compiled binary
    cmds:
      - rm -rf ./trial_task_hello
```

## ✅ Output Examples

```bash
$ task hello
Hello, World from dev

$ task goodbye
Goodbye, cruel world from prod
```

---
"""
