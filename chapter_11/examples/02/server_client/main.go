package main

import (
	"fmt"
	"reflect"

	"example.com/demo/lib/p0"
	"example.com/demo/util"
)

func main() {
	fmt.Print(util.Hello() + "\n")
	fmt.Print("hello 9\n")

	fmt.Println("Hello", p0.Xello())

	intArr := []int{2, 3, 5, 7, 11}

	fmt.Println(reflect.TypeOf(intArr))
	fmt.Println("\n\n", p0.Name)

	//p0.Xello_()
	//p0.SumVals(12, 21)
	p0.UseFunc(p0.SumVals, 12, 21)
}

/*
% go mod init example.com/demo
% go run main.go

*/
