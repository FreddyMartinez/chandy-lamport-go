package models

type Message struct {
	sender   string
	receiver string
	body     interface{}
}

func NewMessage(sender string, receiver string, body interface{}) Message {
	return Message{sender: sender, receiver: receiver, body: body}
}
