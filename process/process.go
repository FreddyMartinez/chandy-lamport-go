package process

import (
	"chandylamport/models"
	"fmt"
)

type Process struct {
	mainJob       *MainJob
	communication *CommunicationModule
	stateManager  *StateManager
}

func CreateProcess(processId int, network []models.ProcessInfo, taskList []models.Task, quit chan bool) *Process {
	processInfo := network[processId]
	fmt.Println(fmt.Sprintf("Creating process: %v, in port: %v", processInfo.Name, processInfo.Port))

	processMessageIn := make(chan models.Message)
	processMessageOut := make(chan models.Message)
	updateStateChan := make(chan models.ProcessEvent) // notify when an event occurs
	saveGlobalState := make(chan bool)

	thisJob := CreateJob(processInfo, network, updateStateChan, processMessageIn, processMessageOut, taskList, quit)

	thisCommunicationMod := CreateCommunicationModule(processInfo.Port, processMessageIn, processMessageOut)

	thisStateManager := CreateStateManager(processId, updateStateChan, saveGlobalState)

	thisProcess := Process{
		mainJob:       thisJob,
		communication: thisCommunicationMod,
		stateManager:  thisStateManager,
	}

	return &thisProcess
}
