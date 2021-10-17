package models

type ProcessData struct {
	amount int // usar mutex?
}

type ProcessEvent struct {
	Description string
	Data        ProcessData // to simulate some task
}

type ProcessHistory struct {
	CurrentEvent int
	EventHistory []ProcessEvent
	MessagesIn   []Message
	MessagesOut  []Message
}

type GlobalState struct {
	NetworkState map[string]ProcessHistory
}

type ChanelState struct {
	Recording bool
}
