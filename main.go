package main

import (
	"chandylamport/helpers"
	"chandylamport/process"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Obtain process data
	args := os.Args[1:]
	processId := args[0]
	fileName := args[1]

	index, err := strconv.Atoi(processId)
	if err != nil {
		panic("Invalid argument when creating process")
	}

	network := helpers.ReadNetConfig(fileName)

	process.CreateProcess(index, network)

	fmt.Println("LLegó al final")
	time.Sleep(50 * time.Second) // usar kill chan bool
}
