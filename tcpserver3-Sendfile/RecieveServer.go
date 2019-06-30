package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		//目录存在
		return true
	}
	if os.IsNotExist(err) {
		//目录不存在
		return false
	}
	return true
}

func main() {
	//0. basic directory ,check if exist,if not ,create one
	dir := "/tmp/uploaddir"
	if !PathExist(dir) {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			//　创建目录失败
			return
		}
	}
	//1. listen
	ln, lnErr := net.Listen("tcp", "127.0.0.1:8888")
	if lnErr != nil {
		fmt.Println("Listen Error:", lnErr)
		return
	}
	defer ln.Close()

	for {
		//2. accept connection
		clientConn, connErr := ln.Accept()
		if connErr != nil {
			fmt.Println("clientConn Error:", connErr)
			return
		}
		go func(conn net.Conn) {
			//读取所有内容，并且存放在dir所指定的文件中
			buf := make([]byte, 1024*4)
			//接收文件名
			// n, connReadErr := conn.Read(buf)
			rd := bufio.NewReader(conn)
			// if connReadErr != nil {
			// 	fmt.Println("Read Client Conn failed", connReadErr)
			// 	return
			// }

			line, _, _ := rd.ReadLine()
			filename := string(line)
			fmt.Println(filename)
			fd, err := os.Create(dir + "/" + filename)
			defer fd.Close()
			if err != nil {
				//创建目标文件失败，退出执行
				fmt.Println("Create File Failed", err)
				return
			}
			for {
				n, err := rd.Read(buf)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("Read client conn failed")
					return
				}
				fd.Write(buf[:n])
			}
		}(clientConn)
	}

}
