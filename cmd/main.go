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

	var chans []chan string

	go handlemsg(chans, chanmsg)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		uchan := make(chan string, 100)
		chans = append(chans, uchan)
		go handleConn(conn, uchan, chanmsg)

	}
}

func handleConn(c net.Conn, uchan chan string, chanmsg chan string) {
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
			close(uchan)
			break
		}
		msg := username + ": " + str
		fmt.Println(msg)
		uchan <- msg
	}

	// Connection closed for some reason
	fmt.Printf("Connection to client %s has been closed\n", username)

}

func handlemsg(uchannels []chan string, msgchan chan string) {

	for _, uchan := range uchannels {
		for {
			select {
			case msg := <-uchan:
				fmt.Printf(msg)
			}
		}
	}
}
