package helpers

import (
	"log"
	"os"

	"github.com/DistributedClocks/GoVector/govec"
)

type Logger struct {
	Event    *log.Logger
	Mark     *log.Logger
	Snapshot *log.Logger
	GoVec    *govec.GoLog
}

func CreateLogger(processId string) *Logger {
	file, err := os.OpenFile("./logs/"+processId+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, processId+"\tEVENT: \t\t", log.Ltime)
	mark := log.New(file, processId+"\tMARK: \t\t", log.Ltime)
	snap := log.New(file, processId+"\tSNAPSHOT: \t", log.Ltime)

	defaultConfig := govec.GetDefaultConfig()
	defaultConfig.UseTimestamps = true
	goVector := govec.InitGoVector(processId, "govector/"+processId, defaultConfig)

	return &Logger{Event: logger, Mark: mark, Snapshot: snap, GoVec: goVector}
}

func (log *Logger) GoVectLog(message string) {
	log.GoVec.LogLocalEvent(message, govec.GetDefaultLogOptions())
}
