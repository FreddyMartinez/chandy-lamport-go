package process

import (
	"chandylamport/helpers"
	"chandylamport/models"
	"fmt"
	"net"
)

type CommunicationModule struct {
	listener          net.Listener
	processMessageOut chan models.Message
	processMessageIn  chan models.Message
}

func CreateCommunicationModule(port string, processMessageIn chan models.Message, processMessageOut chan models.Message) *CommunicationModule {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println(err)
		panic("Server listen error")
	}

	var communicationModule = CommunicationModule{
		listener:          listener,
		processMessageIn:  processMessageIn,
		processMessageOut: processMessageOut,
	}

	go communicationModule.receiver()
	go communicationModule.sender()
	return &communicationModule
}

// This method should wait for messages from other processes
func (comMod *CommunicationModule) receiver() {
	for {
		data := new(models.Message)
		err := helpers.Receive(data, &comMod.listener)
		if err != nil {
			panic(err)
		}
		comMod.processMessageIn <- *data
	}
}

// This method should send any message to other processes, not concurrent
func (comMod *CommunicationModule) sender() {
	for {
		select {
		case processMsg := <-comMod.processMessageOut:
			// find ip and port first
			// helpers.Send(processMsg, processMsg.Receiver)
			fmt.Println(processMsg)
		}
	}
}
