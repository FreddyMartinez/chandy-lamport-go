package main_test

import (
	"chandylamport/helpers"
	"os"
	"strconv"
	"testing"
	"time"
)

const localPath = "/home/freedy/Documents/master/RedesYDistribuidos/Practica1/chandylamport/"

// Used to launch all processes programmatically
func TestMain(m *testing.M) {
	// crear los 3 procesos
	for i := 0; i < 3; i++ {
		sshConn := helpers.CreateSSHClient("127.0.0.1")

		command := "cd " + localPath + "; /usr/local/go/bin/go run main.go " + strconv.Itoa(i) + " network.json"
		go helpers.RunCommand(command, sshConn)

		defer sshConn.Close()
	}
	time.Sleep(10 * time.Second)
	code := m.Run()
	os.Exit(code)
}
