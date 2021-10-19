package models

type Task struct {
	MsgType      string
	Time         int
	Receiver     int
	Value        int
	NetworkDelay int
}
