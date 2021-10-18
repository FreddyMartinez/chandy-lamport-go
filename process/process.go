package process

import (
	"chandylamport/models"
	"fmt"
)

type Process struct {
	mainJob          *models.MainJob
	communication    *CommunicationModule
	stateManager     StateManager
	processMessageIn chan models.Message
}

func CreateProcess(processId int, network []models.ProcessInfo) *Process {
	processInfo := network[processId]
	fmt.Println(fmt.Sprintf("Creating process: %v, in port: %v", processInfo.Name, processInfo.Port))

	processMessageIn := make(chan models.Message)
	processMessageOut := make(chan models.Message)
	updateStateChan := make(chan models.ProcessEvent) // notify when an event occurs

	thisJob := models.CreateJob(processInfo, network, updateStateChan, processMessageIn, processMessageOut)

	thisCommunicationMod := CreateCommunicationModule(processInfo.Port, processMessageIn, processMessageOut)

	thisProcess := Process{
		mainJob:          thisJob,
		communication:    thisCommunicationMod,
		processMessageIn: processMessageIn,
	}

	go thisProcess.ReceiveMessages()
	return &thisProcess
}

// Receive mesages from Communication Module
func (p *Process) ReceiveMessages() {
	fmt.Println("Lanza ReceiveMessages")
	for {
		message := <-p.processMessageIn
		fmt.Printf("Msg [%v] receive from: %s\n", message.Body, message.Sender)
	}
}
