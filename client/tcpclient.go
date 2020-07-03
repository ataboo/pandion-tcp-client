package client

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

type TcpClient struct {
	conn     net.Conn
	lock     sync.Mutex
	stopChan chan int
}

func NewTcpClient() *TcpClient {
	return &TcpClient{
		lock: sync.Mutex{},
	}
}

func (c *TcpClient) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	c.conn = conn

	return err
}

func (c *TcpClient) Send(message string) error {
	if c.conn == nil {
		return fmt.Errorf("No connection started")
	}
	_, err := c.conn.Write([]byte(message))

	return err
}

func (c *TcpClient) Read() (string, error) {
	if c.conn == nil {
		return "", fmt.Errorf("No connection started")
	}
	reader := bufio.NewReader(c.conn)
	if reader.Buffered() == 0 {
		return "", fmt.Errorf("Nothing to read from connection")
	}
	str, err := reader.ReadString('\n')

	return str, err
}

func (c *TcpClient) StartReadPump(handleLine func(line string)) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.stopChan != nil {
		return fmt.Errorf("Read pump already started.")
	}

	c.stopChan = make(chan int)
	tickChan := time.Tick(time.Second)
	go func() {
		for {
			select {
			case <-c.stopChan:
				return
			case <-tickChan:
				line, err := c.Read()
				if err != nil {
					handleLine(line)
				}
			}
		}
	}()

	return nil
}

func (c *TcpClient) Close() error {
	if c.conn == nil {
		return fmt.Errorf("No connection to close")
	}

	c.conn.Close()

	return nil
}
