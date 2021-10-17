package main

import (
	"chandylamport/models"
	"chandylamport/process"
	"fmt"
	"os"
	"time"
)

func main() {
	// Obtain process data
	args := os.Args[1:]
	processId := args[0]
	port := args[1]
	ip := args[2]
	fmt.Println(fmt.Sprintf("Creating process: %v, in port: %v", processId, port))

	processInfo := models.ProcessInfo{Port: port, Name: processId, Ip: ip}

	process := process.CreateProcess(processInfo)

	go process.ReceiveMessages()

	fmt.Println("LLegó al final")
	time.Sleep(50 * time.Second)
}
