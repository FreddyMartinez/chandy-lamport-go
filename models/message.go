package models

type Message struct {
	Sender   string
	Receiver string
	Body     interface{}
}

func NewMessage(sender string, receiver string, body interface{}) Message {
	return Message{Sender: sender, Receiver: receiver, Body: body}
}
