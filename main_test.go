package main_test

import (
	"chandylamport/helpers"
	"os"
	"strconv"
	"testing"
	"time"
)

// const localPath = "/home/a846866/Documents/"
const localPath = "/home/freedy/Documents/master/RedesYDistribuidos/Practica1/chandylamport/"

// const binFile = "; ./chandylamportlab "
const binFile = "; ./chandylamport "

// const jsonFile = " labnetwork.json"
const jsonFile = "network.json"

// Used to launch all processes programmatically
func TestMain(m *testing.M) {
	// crear los 3 procesos
	processes := helpers.ReadNetConfig(jsonFile)
	for i, proc := range processes {
		sshConn := helpers.CreateSSHClient(proc.Ip)

		command := "cd " + localPath + binFile + strconv.Itoa(i) + " " + jsonFile
		go helpers.RunCommand(command, sshConn)

		defer sshConn.Close()
	}
	time.Sleep(10 * time.Second)
	code := m.Run()
	os.Exit(code)
}
