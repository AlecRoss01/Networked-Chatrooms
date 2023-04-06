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

type User struct {
	Username string
}

func makeConn() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println(err)
	}
	return conn
}

func joinChatroom() {
	c := makeConn()
	data := make(chan string, 100)
	exit := make(chan string)
	go readConn(c, exit, data)
	fmt.Println("Welcome to the chatroom!")
	for {
		select {
		case message := <-data:
			fmt.Print(message)
		printall:
			for {
				select {
				case message := <-data:
					fmt.Print(message)
				default:
					break printall
				}
			}

		default:
		}
		input := readInput()
		val, err := c.Write([]byte(input))
		if err != nil {
			fmt.Println(err)
			fmt.Print(val)
			return
		} else if input == "!exit\n" {
			exit <- "done"
			return
		}
	}
}

func readConn(c net.Conn, exit chan string, data chan string) {
	reader := bufio.NewReader(c)
	message, _ := reader.ReadString('\n')
	fmt.Print(message)
	if message != "\n" {
		//fmt.Print(message)
		data <- message
	}
	select {
	case <-exit:
		c.Close()
		return
	default:
	}

}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
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

func commandHandler(command string, users []User) string {
	echo := "!echo "
	listUsers := "!users"
	if strings.Contains(command, echo) {
		location := strings.Index(command, echo)
		return command[location+len(echo):] + "\n"
	}
	if strings.Contains(command, listUsers) {
		return formatUserList(users)
	}

	return ""
}

func formatUserList(users []User) string {
	var sb strings.Builder
	sb.WriteString("Users in Chatroom:\n")
	for _, user := range users {
		sb.WriteString(user.Username)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func main() {
	joinChatroom()
}
