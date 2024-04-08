package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
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

	defer conn.Close()

	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err2 = ", err2)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))

}
