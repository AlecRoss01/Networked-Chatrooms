package main

import (
	"bufio"
	"fmt"
	"net"
)

func tcpHandler() {
	formatPort := ":30000"
	l, err := net.Listen("tcp", formatPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	readConn(c)
	c.Write([]byte("Hello!\n"))
	c.Close()
}

func readConn(c net.Conn) {
	reader := bufio.NewReader(c)
	message, _ := reader.ReadString('\n')
	fmt.Println(message)
}

func main() {
	tcpHandler()
}
