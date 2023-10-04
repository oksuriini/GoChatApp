package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":30222")
	chanmsg := make(chan string)
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
		go handleConn(conn, uchan, chanmsg)

	}
}

func handleConn(c net.Conn, uchan chan string, chanmsg chan string) {
	reader := bufio.NewReader(c)
	c.Write([]byte("Please enter your name: "))
	username, err := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	go handlemsg(chanmsg)
	if err != nil {
		fmt.Println("Error creating username")
	}

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("SOMETHING WENT TERRIBLY WRONG")
			close(uchan)
			break
		}
		msg := username + ": " + strings.TrimSpace(str)
		chanmsg <- msg
	}

	// Connection closed for some reason
	fmt.Printf("Connection to client %s has been closed\n", username)

}

func handlemsg(msgchan chan string) {
	for {
		select {
		case msg := <-msgchan:
			fmt.Println(msg)
		}
	}
}
