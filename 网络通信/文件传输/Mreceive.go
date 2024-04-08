package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(conn net.Conn, filename string) {
	//新建文件
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完成")
			} else {
				fmt.Println("err = ", err)
			}
			return
		}
		f.Write(buf[:n])
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//第一次读取，保存文件名
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("err3 = ", err3)
		return
	}
	fileName := string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//接收文件
	RecvFile(conn, fileName)

}
