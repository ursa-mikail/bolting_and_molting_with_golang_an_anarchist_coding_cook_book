package main

import (
    "fmt"
    "reflect"

    "test-app/lib/p0"
    "test-app/util"
)

func main() {
    fmt.Print(util.Hello() + "\n")
    fmt.Print("hello 9\n")

    fmt.Println("Hello", p0.Xello())

    intArr := []int{2, 3, 5, 7, 11}

    fmt.Println(reflect.TypeOf(intArr))
    fmt.Println("\n\n", p0.Name)

    p0.UseFunc(p0.SumVals, 12, 21)
}

/*
% go mod init test-app
% go run main.go
*/
