package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		println("Expected a single address argument")
		os.Exit(1)
	}

	addr := args[1]
	readBuffer := make([]byte, 1024)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		println("Failed to resolve address", err)
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		println("Failed to listen on port", err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			println("Failed to accept conn", err)
			os.Exit(1)
		}

		n, err := conn.Read(readBuffer)
		if err != nil {
			log.Fatal("Failed to read", err)
			return
		}

		fmt.Printf("Got line: %s\n", string(readBuffer[:n]))

		if string(readBuffer[:n]) == "ping" {
			_, err := conn.Write([]byte("pong"))
			if err != nil {
				fmt.Printf("Failed to write echo: %s", err)
			}
		} else if string(readBuffer[:n]) == "ping2" {
			_, err := conn.Write([]byte("pong1\n"))
			if err != nil {
				fmt.Printf("Failed to write echo: %s\n", err)
			}

			_, err = conn.Write([]byte("pong2\n"))
			if err != nil {
				fmt.Printf("Failed to write echo: %s\n", err)
			}
		}
	}
}
