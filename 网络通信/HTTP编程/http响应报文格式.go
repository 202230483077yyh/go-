package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer conn.Close()

	//请求包报文
	requestBuf := ""

	//宪法请求包，服务器才会回响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的包
	buf := make([]byte, 1024*4)

	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("read err1 = ", err1)
		return
	}

	//打印响应报文
	fmt.Printf("#%v#", string(buf))

}
