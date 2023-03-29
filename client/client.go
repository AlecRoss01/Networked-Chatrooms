package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	//"strings"
)
func makeConn() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println(err)
	return
	}
	val, err := conn.Write([]byte(readInput()));
	if err != nil {
		fmt.Println(err)
		fmt.Print(val)
	return
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the chatroom!")
	text, _ := reader.ReadString('\n')
	//text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	makeConn()
}
