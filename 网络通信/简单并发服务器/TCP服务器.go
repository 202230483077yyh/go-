package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完毕自动关闭conn
	defer conn.Close()

	//获取网络端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful")

	buf := make([]byte, 2048)

	for {
		//读取用户数据
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			return
		}

		//fmt.Println("服务器读取长度为：", n)
		//之所以是n-1，因为换行多了1个换行字符
		str := string(buf[:n-1])
		if str == "exit" {
			fmt.Println("addr exit")
			break
		}
		fmt.Printf("[%s]: %s\n ", addr, string(buf[:n]))
		//把数据转化为大写再发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()
	//接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		//处理用户请求,新建一个协程
		go HandleConn(conn)
	}
}
