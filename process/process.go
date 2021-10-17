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

func CreateProcess(processInfo models.ProcessInfo) *Process {

	processMessageIn := make(chan models.Message, 10)
	processMessageOut := make(chan models.Message)
	updateStateChan := make(chan models.ProcessEvent) // notify when an event occurs

	// Obtener de un archivo
	network := []models.ProcessInfo{
		{
			Name: "P0",
			Ip:   "127.0.0.1",
			Port: "18660",
		},
		{
			Name: "P1",
			Ip:   "127.0.0.1",
			Port: "18661",
		},
	}

	thisJob := models.CreateJob(processInfo, network, updateStateChan, processMessageOut)

	thisCommunicationMod := CreateCommunicationModule(processInfo.Port, processMessageIn, processMessageOut)

	thisProcess := Process{
		mainJob:          thisJob,
		communication:    thisCommunicationMod,
		processMessageIn: processMessageIn,
	}

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
