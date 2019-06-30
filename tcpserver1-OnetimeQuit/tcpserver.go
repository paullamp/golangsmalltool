package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. listen in port 127.0.0.1:8888
	ln, lnerr := net.Listen("tcp", "127.0.0.1:8888")
	if lnerr != nil {
		fmt.Println("listen and bind error:", lnerr)
		return
	}
	// 2. waiting for client connections
	conn, connerr := ln.Accept()
	if connerr != nil {
		fmt.Println("Accept client connection failed")
		return
	}

	// 3. send something to client
	_, writeerr := conn.Write([]byte("Hello , this is Server"))
	if writeerr != nil {
		fmt.Println("Write to client Error:", writeerr)
		return
	}
	// 4. receive something from client
	buf := make([]byte, 1024)
	n, connReaderr := conn.Read(buf)
	if connReaderr != nil {
		fmt.Println("connRead Error:", connReaderr)
		return
	}
	fmt.Println(string(buf[:n]))
}
