package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func makeConn() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := conn.Write([]byte(readInput()))
	if err != nil {
		fmt.Println(err)
		fmt.Print(val)
		return
	}
	readConn(conn)
}

func readConn(c net.Conn) {
	reader := bufio.NewReader(c)
	message, _ := reader.ReadString('\n')
	fmt.Println(message)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the chatroom!")
	text, _ := reader.ReadString('\n')
	//text = strings.Replace(text, "\n", "", -1)
	return text
}

func thirdPerson(text string, name string) string {
	return name + " " + text
}

func coinFlip() string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 1
	result := rand.Intn(2)

	// If the result is 0, return "heads", otherwise return "tails"
	if result == 0 {
		return "heads"
	} else {
		return "tails"
	}
}

func diceRoll() int {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	result := rand.Intn(6) + 1

	return result
}

func commandHandler(command string) string {
	echo := "!echo "
	if strings.Contains(command, echo) {
		location := strings.Index(command, echo)
		return command[location+len(echo):] + "\n"
	}
	return ""
}

func main() {
	makeConn()
}
