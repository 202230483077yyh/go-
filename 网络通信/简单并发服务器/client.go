package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//main调用完毕，关闭链接
	defer conn.Close()

	//接受服务器回复的数据，新任务
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	//从键盘输入内容，给服务器发送
	str := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//fmt.Println("客户端传送长度n为 ；", n)
		//把输入的内容发送给服务器
		conn.Write(str[:n])
		if string(str[:n-1]) == "exit" {
			fmt.Println("client exit")
			return
		}

	}

}
