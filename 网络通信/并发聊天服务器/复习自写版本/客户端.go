package main

import (
	"fmt"
	"net"
)

func readconn(conn net.Conn) {
	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err1 = ", err)
			return
		}
		str := string(buf[:n])
		fmt.Println(str)
	}
}

func writeconn(conn net.Conn) {
	var str string
	for {
		fmt.Scan(&str)
		conn.Write([]byte(str))
		if str == "exit" {
			fmt.Println("客户端主动断开链接")
			return
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	go readconn(conn)
	//为了保证输入exit之后全部结束，readconn不起作用，exit必须放在主协程一级
	writeconn(conn)

}
