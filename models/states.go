package models

import "sync"

type ProcessData struct {
	Amount int
	Mu     sync.Mutex
}

type ProcessEvent struct {
	Description string
	Data        string // to simulate some task
}

type ProcessHistory struct {
	CurrentEvent int
	EventHistory []ProcessEvent
}

type Snapshot struct {
	EventHistory ProcessHistory
	MessagesIn   []Message
	MessagesOut  []Message
}

type ChanelState struct {
	Recording bool
}
