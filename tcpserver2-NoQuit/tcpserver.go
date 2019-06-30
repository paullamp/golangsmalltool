package main

import (
	"fmt"
	"net"
	// "time"
)

func main() {
	// 1. Listen
	ln, lnErr := net.Listen("tcp", "127.0.0.1:8888")
	if lnErr != nil {
		fmt.Println("Listen error:", lnErr)
		return
	}

	//2. Accept connections

	for {
		conn, connErr := ln.Accept()
		if connErr != nil {
			fmt.Println("connection Err:", connErr)
			continue
		}

		go func(conn net.Conn) {
			// conn.Write([]byte("Hello Client"))
			conn.Write([]byte(time.Now().String() + ": Hello Client"))
			// time.Sleep(10 * time.Second)
			//wait for 10 second,to test if the server can handle more connections
			conn.Close()
		}(conn)
	}
}
