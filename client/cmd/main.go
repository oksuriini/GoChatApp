package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var addr *string
var port *string
var username *string

func init() {
	addr = flag.String("Address", "127.0.0.1", "Connect to this specific address")
	port = flag.String("Port", ":9090", "Connect to this specific port of the address")
	username = flag.String("Username", "Anonymous", "Set username")
}

func main() {

	println(*addr + *port)

	// conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", *addr, *port))
	conn, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		fmt.Println("Something went wrong trying to establish connection")
		os.Exit(2)
	}
	write := bufio.NewWriter(conn)
	for {
		_, err := write.ReadFrom(os.Stdin)
		if err != nil {
			fmt.Println("Wroong")
			return
		}
	}

}
