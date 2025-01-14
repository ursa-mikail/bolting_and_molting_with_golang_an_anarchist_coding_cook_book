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
