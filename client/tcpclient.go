package client

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

// TCPClient wraps net's tcpclient
type TCPClient struct {
	address string
	conn    net.Conn
	buffer  *bytes.Buffer
}

// NewTCPClient create a new TCPClient
func NewTCPClient(address string) *TCPClient {
	return &TCPClient{
		address: address,
		conn:    nil,
		buffer:  bytes.NewBuffer(make([]byte, 1024)),
	}
}

// Connect dial the tcp connection.
func (c *TCPClient) Connect() error {
	conn, err := net.Dial("tcp", c.address)
	c.conn = conn

	return err
}

// ReadString Read a string from the connection.
func (c *TCPClient) ReadString() (str string, err error) {
	if c.conn == nil {
		return "", fmt.Errorf("No connection started")
	}

	tempBuffer := make([]byte, 1024)

	c.buffer.Reset()
	c.conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, err := c.conn.Read(tempBuffer)
	if err != nil {
		return "", err
	}

	fmt.Printf("Read %d bytes", n)

	return string(tempBuffer[:n]), nil
}

//WriteString write a string to the connection.
func (c *TCPClient) WriteString(message string) error {
	if c.conn == nil {
		return fmt.Errorf("No connection started")
	}

	c.conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	_, err := c.conn.Write([]byte(message))

	return err
}

// Close Close the connection.
func (c *TCPClient) Close() error {
	if c.conn == nil {
		return fmt.Errorf("No connection to close")
	}

	c.conn.Close()
	c.conn = nil

	return nil
}
