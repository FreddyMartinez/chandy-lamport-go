package models

type ProcessInfo struct {
	Name string
	Ip   string
	Port string
}

type MainJob struct {
	ProcessInfo ProcessInfo   // it's own info
	NetworkInfo []ProcessInfo // other processes info
	Data        ProcessData
}

func CreateJob(processInfo ProcessInfo) *MainJob {
	myJob := MainJob{ProcessInfo: processInfo}
	return &myJob
}
