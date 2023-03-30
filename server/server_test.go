package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func init() {
	// Run the servers in goroutines to stop blocking
	go func() {
		tcpHandler()
	}()
}

func TestNETServer_Run(t *testing.T) {
	// Simply check that the server is up and can
	// accept connections.
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	val, err := conn.Write([]byte("Hello\n"))
	if err != nil {
		fmt.Printf("err: %q, val:, %q", err, val)
		return
	}
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		t.Error("connection error", err)
	}
	if message != "Hello!\n" {
		t.Errorf("message was not what was expected, got %q, wanted \"Hello!\\\n\"", message)
	}
	defer conn.Close()
	conn.Close()
}
