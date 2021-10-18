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
	processInfo := network[index]

	fmt.Println(fmt.Sprintf("Creating process: %v, in port: %v", processInfo.Name, processInfo.Port))

	process := process.CreateProcess(processInfo)

	go process.ReceiveMessages()

	fmt.Println("LLegó al final")
	time.Sleep(5 * time.Second)
}
