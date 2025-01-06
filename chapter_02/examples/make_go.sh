# Function to print usage and exit
usage() {
    echo "Usage: $0 <module-name>"
    echo "Example: $0 example.com/demo"
    exit 1
}

# Check if module name is provided
if [ $# -ne 1 ]; then
    usage
fi

MODULE_NAME=$1

# Create project structure
mkdir -p "$MODULE_NAME/libs/p0" "$MODULE_NAME/util"

# Create go.mod
cd "$MODULE_NAME" || exit
go mod init "$MODULE_NAME"

# Create libs/p0/p0.go
cat > libs/p0/p0.go <<EOF
package p0

import "fmt"

var Name = "DefaultName"

func Xello() string {
    return "World"
}

func SumVals(a, b int) int {
    return a + b
}

func UseFunc(f func(int, int) int, a, b int) {
    result := f(a, b)
    fmt.Println("Result:", result)
}
EOF

# Create utils/util_00.go
cat > utils/util_00.go <<EOF
package utils

func Hello() string {
    return "Hello from util"
}
EOF

# Create main.go
cat > main.go <<EOF
package main

import (
    "fmt"
    "reflect"

    "$MODULE_NAME/libs/p0"
    "$MODULE_NAME/utils"
)

func main() {
    fmt.Print(utils.Hello() + "\\n")
    fmt.Print("hello 9\\n")

    fmt.Println("Hello", p0.Xello())

    intArr := []int{2, 3, 5, 7, 11}

    fmt.Println(reflect.TypeOf(intArr))
    fmt.Println("\\n\\n", p0.Name)

    p0.UseFunc(p0.SumVals, 12, 21)
}

/*
% go mod init $MODULE_NAME
% go run main.go
*/
EOF

echo "Go project setup complete! Navigate to '$MODULE_NAME' and start coding."

