package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtain process data
	args := os.Args[1:]
	processId := args[0]
	port := args[1]
	fmt.Println(fmt.Sprintf("Creating process: %v, in port: %v", processId, port))
}
