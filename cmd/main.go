package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":30222")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		uchan := make(chan string, 100)
		go handleConn(conn, uchan)
	}
}

func handleConn(c net.Conn, cha chan string) {
	reader := bufio.NewReader(c)
	c.Write([]byte("Please enter your name: "))
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		fmt.Println("Error creating username")
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("SOMETHING WENT TERRIBLY WRONG")
			break
		}
		msg := username + ": " + str
		c.Write([]byte(msg))
	}

	// Connection closed for some reason
	fmt.Printf("Connection to client %s has been closed\n", username)

}
