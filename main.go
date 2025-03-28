package main

import (
	"fmt"
	"net"
)

func main() {
	// Start a TCP listener:
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port :6379")

	// Start receiving requests:
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close() // close connection once finished

	// Create an infinite loop and receive commands from clients and respond to them:
	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		_ = value

		writer := NewWriter(conn)
		writer.Write(Value{typ: "string", str: "OK"})
	}
}
