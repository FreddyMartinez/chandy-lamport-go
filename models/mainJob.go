package models

type ProcessInfo struct {
	name string
	ip   string
	port string
}

type MainJob struct {
	processInfo ProcessInfo   // it's own info
	networkInfo []ProcessInfo // other processes info
	data        ProcessData
}

func CreateJob() *MainJob {
	myJob := MainJob{}
	return &myJob
}
