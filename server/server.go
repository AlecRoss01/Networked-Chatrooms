package main

import (
	"bufio"
	"fmt"
	"net"
)

func enterChatroom(c net.Conn, ch chan string, mainCh chan string) {
	readChan := make(chan string, 10)
	go readConn(c, readChan)
	for {
		select {
			// get data from readConn and send it to main
		case data := <- readChan:
			if data == "!exit" {
				return
			} else {
				mainCh <- data
			}
			// get data to be sent to client and send
		case data := <- ch:
			c.Write([]byte(data))
		}
	}
	//data := readConn(c)
	//mainCh <- data
	//c.Write([]byte("Hello!\n"))
	//c.Close()
}

func sendToChatroom(chans []chan string, data string) {
	for i := 0; i < len(chans); i++ {
		chans[i] <- data
	}
}

func tcpHandler() {
	main := make(chan string, 100)
	var chans []chan string
	formatPort := ":30000"
	l, err := net.Listen("tcp", formatPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		newCh := make(chan string, 10)
		chans = append(chans, newCh)
		go enterChatroom(c, newCh, main)
	}
}

func readConn(c net.Conn, ch chan string) {
	reader := bufio.NewReader(c)
	message, _ := reader.ReadString('\n')
	//fmt.Println(message)
	ch <- message
	//return message
}

func main() {
	tcpHandler()
}
