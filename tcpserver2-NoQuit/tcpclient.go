package main

import (
	"fmt"
	"net"
)

func main() {
	//1. connect to server
	conn, connErr := net.Dial("tcp", "127.0.0.1:8888")
	if connErr != nil {
		fmt.Println("connect to server error:", connErr)
		return
	}
	defer conn.Close()
	//2. read from server
	buf := make([]byte, 1024)
	n, readErr := conn.Read(buf)
	if readErr != nil {
		fmt.Println("Read Error:", readErr)
		return
	}
	fmt.Println(string(buf[:n]))
}
