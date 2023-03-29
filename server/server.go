package main

import (
	"fmt"
	"net"
	"bufio"
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
	reader := bufio.NewReader(c)
	message, _ := reader.ReadString('\n')
	fmt.Println(message)
	c.Write([]byte("Hello!"))
}

func main() {
	tcpHandler()
}