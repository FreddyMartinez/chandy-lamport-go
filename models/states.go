package models

type ProcessData struct {
	amount int // usar mutex?
}

type ProcessEvent struct {
	description string
	data        ProcessData // to simulate some task
}

type ProcessHistory struct {
	currentEvent int
	eventHistory []ProcessEvent
	messagesIn   []Message
	messagesOut  []Message
}

type GlobalState struct {
	networkState map[string]ProcessHistory
}

type ChanelState struct {
	recording bool
}
