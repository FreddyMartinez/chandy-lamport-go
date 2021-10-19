package models

const MsgApp = "AppMsg"
const MsgMark = "MarkMsg"

type Message struct {
	Sender       string
	Receiver     string
	MsgType      string
	NetworkDelay int
	Body         int
}

func NewAppMessage(sender string, receiver string, body int, networkDelay int) Message {
	return Message{Sender: sender, Receiver: receiver, MsgType: MsgApp, Body: body, NetworkDelay: networkDelay}
}

func NewMarkMessage(sender string, receiver string, networkDelay int) Message {
	return Message{Sender: sender, Receiver: receiver, MsgType: MsgMark, NetworkDelay: networkDelay}
}
