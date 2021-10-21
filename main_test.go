package main_test

import (
	"chandylamport/helpers"
	"os"
	"strconv"
	"testing"

	"golang.org/x/crypto/ssh"
)

const localPath = "/home/freedy/Documents/master/RedesYDistribuidos/Practica1/chandylamport/"

// Used to launch all processes programmatically
func MainTest(m *testing.M) {
	// crear los 3 procesos
	sshConn := make(map[int]*ssh.Client)
	for i := 0; i < 3; i++ {
		sshConn[i] = helpers.CreateSSHClient("127.0.0.1")

		command := "/usr/local/go/bin/go run " + localPath + "main.go " + strconv.Itoa(i) + " network.json"
		go helpers.RunCommand(command, sshConn[i])

	}
	for _, conn := range sshConn {
		defer conn.Close()
	}
	code := m.Run()
	os.Exit(code)
}
