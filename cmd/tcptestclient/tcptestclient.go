package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	readBuffer := make([]byte, 1024)

	conn, err := net.Dial("tcp", "127.0.0.1:3001")
	if err != nil {
		log.Fatal(err)
	}

	println("Connected!")

	defer conn.Close()

	_, err = conn.Write([]byte("blah!"))
	if err != nil {
		log.Fatal("Failed to send!")
	}

	_, err = conn.Read(readBuffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got bytes: %s\n", string(readBuffer))
}
