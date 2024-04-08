package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.open err:", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	//读取文件内容
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.read err = ", err)
			}
			return
		}
		conn.Write(buf[:n])
	}

}

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件名： ")
	var path string
	fmt.Scan(&path)

	//获取文件名info.Name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//主动链接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	defer conn.Close()

	//给接收方发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//接收对方回复
	buf := make([]byte, 1024)
	m, err3 := conn.Read(buf)

	if err3 != nil {
		fmt.Println("err3 = ", err3)
		return
	}

	if "ok" == string(buf[:m]) {
		//发送文件内容
		sendFile(path, conn)
	}

}
