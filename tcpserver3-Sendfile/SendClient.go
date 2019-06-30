package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//获取输入的参数
	if len(os.Args) < 2 {
		fmt.Println("Please input the file need to send.")
		fmt.Println("Usage: SendClient srcfilename")
		return
	}
	//1. connect to server
	conn, connErr := net.Dial("tcp", "127.0.0.1:8888")
	if connErr != nil {
		fmt.Println("connect to Server error:", connErr)
		return
	}
	defer conn.Close()

	//2. open and readfile
	srcfile := os.Args[1]
	fd, fdErr := os.Open(srcfile)
	if fdErr != nil {
		fmt.Println("Open src file failed")
		return
	}
	defer fd.Close()
	//3. read file and send
	buf := make([]byte, 1024*4)
	//传输文件名至对端服务器
	conn.Write([]byte(srcfile + "\n"))

	for {
		n, fdErr := fd.Read(buf)
		if fdErr == io.EOF {
			break
		}
		if fdErr != nil {
			fmt.Println("read file failed")
			return
		}
		_, writeErr := conn.Write(buf[:n])
		if writeErr != nil {
			fmt.Println("Write to server Error:", writeErr)
		}
	}

}
