package process

import (
	"chandylamport/models"
	"fmt"
)

type Process struct {
	mainJob           *models.MainJob
	communication     *CommunicationModule
	stateManager      StateManager
	updateStateChan   chan models.ProcessEvent // notify when an event occurs
	processMessageOut chan models.Message
	processMessageIn  chan models.Message
}

func CreateProcess(processInfo models.ProcessInfo) *Process {

	processMessageIn := make(chan models.Message, 10)
	processMessageOut := make(chan models.Message)

	thisJob := models.CreateJob(processInfo)

	thisCommunicationMod := CreateCommunicationModule(processInfo.Port, processMessageIn, processMessageOut)

	thisProcess := Process{
		mainJob:           thisJob,
		communication:     thisCommunicationMod,
		processMessageOut: processMessageOut,
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

func (p *Process) SendMessage(message *models.Message) {
	fmt.Printf("Msg [%v] sent to: %s\n", message.Body, message.Receiver)
	fmt.Println(message)
	p.processMessageOut <- *message
}
