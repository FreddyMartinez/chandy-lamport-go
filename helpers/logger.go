package helpers

import (
	"log"
	"os"
)

func CreateLogger(processId string) *log.Logger {
	file, err := os.OpenFile("./logs/"+processId+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, processId+":\t", log.Ltime)

	return logger
}
