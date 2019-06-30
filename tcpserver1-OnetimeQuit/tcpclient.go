package main

import (
	"fmt"
	"net"
)

func main() {
	//1.connect to server
	conn, dialErr := net.Dial("tcp", "127.0.0.1:8888")
	if dialErr != nil {
		fmt.Println("Dial to tcp server failed.", dialErr)
		return
	}
	defer conn.Close()
	//2. read from conn
	buf := make([]byte, 1024)
	n, readErr := conn.Read(buf)
	if readErr != nil {
		fmt.Println("Read Error:", readErr)
		return
	}
	fmt.Println(string(buf[:n]))

	//3. send message to server
	_, writeErr := conn.Write([]byte("Hello Server"))
	if writeErr != nil {
		fmt.Println("Write Error:", writeErr)
		return
	}
}
