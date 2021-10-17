// This module contains the communication functions
// Taken from class slides
package helpers

import (
	"encoding/gob"
	"net"
)

// Send any data to desired ip
func Send(data interface{}, ip string) error {
	var conn net.Conn
	var err error
	var encoder *gob.Encoder

	conn, err = net.Dial("tcp", ip)

	if err != nil {
		panic("Client connection error")
	}

	encoder = gob.NewEncoder(conn)
	err = encoder.Encode(data)
	conn.Close()
	return err
}

// listen for messages
func Receive(data interface{}, listener *net.Listener) error {
	var conn net.Conn
	var err error
	var decoder *gob.Decoder

	conn, err = (*listener).Accept()
	if err != nil {
		panic("Server accept connection error")
	}

	decoder = gob.NewDecoder(conn)

	err = decoder.Decode(data)
	conn.Close()

	return err
}
