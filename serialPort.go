package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

/*
	info: open serial port with standard settings. And return a pointer to the port
	requers: string containing the port name.
	returns: pointer to opend port
*/
func OpenPort(selPort string, baud uint32) *serial.Port {
	port := "/dev/" + selPort
	c := &serial.Config{
		Name:        port,
		Baud:        int(baud),
		ReadTimeout: -1, //don't stop reading the port
	}

	connection, err := serial.OpenPort(c)
	fatalErrorDisplay(err)
	connection.Flush()
	log.Println("Opened port: ", selPort)
	return connection
}

/*
	info: close the opened port
	requers: pointer to serial port
*/
func closePort(stream *serial.Port) {
	stream.Close()
	fmt.Println("port clossed")
}
