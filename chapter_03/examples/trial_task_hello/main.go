package main

import (
	"flag"
	"fmt"
)

var (
	mode        string
	environment string
)

func main() {
	flag.StringVar(&mode, "mode", "hello", "Mode of operation: hello | goodbye")
	flag.StringVar(&environment, "env", "dev", "Environment: dev | prod")
	flag.Parse()

	switch mode {
	case "hello":
		runHello()
	case "goodbye":
		runGoodbye()
	default:
		fmt.Println("Unknown mode:", mode)
	}
}

func runHello() {
	fmt.Println("Hello, World from", environment)
}

func runGoodbye() {
	fmt.Println("Goodbye, cruel world from", environment)
}

/*
task hello
task goodbye
task hello:prod
task goodbye:dev
task clear_executable
*/
