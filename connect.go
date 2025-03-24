package main

import (
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

func connect() int {
	log.Println("Serial Connect")

	mode := &serial.Mode{
		BaudRate: BaudRate,
	}

	port, err := serial.Open(FileDescriptor, mode)
	port.SetReadTimeout(time.Millisecond)

	if err != nil {
		port.Close()

		log.Fatal("serial open err: ", err)
	}

	write_data := make([]byte, 0, 32)
	write_data = append(write_data, "$CC"...)
	checksum := calculate_checksum(write_data)
	write_data = append(write_data, '*')
	write_data = append(write_data, checksum...)
	write_data = append(write_data, '\r')
	write_data = append(write_data, '\n')

	log.Println("writing: ", string(write_data))

	n, err := port.Write(write_data)
	if err != nil {
		log.Fatal("write error: ", err)
	}

	read_data := make([]byte, 32)
	n, err = port.Read(read_data)
	if err != nil {
		log.Fatal("read error: ", err)
	}
	if n == 0 {
		log.Fatal("response: none")
	}

	port.Close()
	read_string := string(read_data)
	log.Printf("received: %v", read_string)

	// check response
	if !strings.HasPrefix(read_string, "$CR") {
		log.Fatal("response: incorrect")
	}

	return 0
}
