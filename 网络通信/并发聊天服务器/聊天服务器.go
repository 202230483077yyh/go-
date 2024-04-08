package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //	发送数据管道
	Name string      //用户名
	Addr string      //网络地址
}

// 保存在线用户 	Addr --->Client

var onlinemap map[string]Client
var message = make(chan string)

// 每一个客户端发送任务
func WriteMsgClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

// 创建任务对象并存入map，同时再开子协程、广播
func HandleConn(conn net.Conn) {
	//获取客户端网络地址
	CliAddr := conn.RemoteAddr().String()

	//创建一个结构体
	cli := Client{make(chan string), CliAddr, CliAddr}

	//把结构体添加到map
	onlinemap[CliAddr] = cli

	//开一个协程给当前客户端发送信息
	go WriteMsgClient(cli, conn)

	message <- "[" + cli.Addr + "] " + cli.Name + " login"

	//对方是否主动退出
	isQuit := make(chan bool)
	hasdata := make(chan bool)

	//新开一个协程接收用户发送数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开或者出问题
				isQuit <- true
				fmt.Println("conn.Read.err = ", err)
				return
			}
			msg := string(buf[:n-1])

			if msg == "who" {
				conn.Write([]byte("user list :\n"))
				for _, val := range onlinemap {
					conn.Write([]byte(val.Name))
				}
			} else if len(msg) > 9 && msg[:6] == "rename" {
				nname := strings.Split(msg, "|")[1]
				cli.Name = nname
				conn.Write([]byte("名字修改成功！"))
			} else {
				message <- msg
			}
			hasdata <- true
		}
	}()

	for {
		//通过select检测channel的流动
		select {
		case <-isQuit:
			delete(onlinemap, CliAddr)
			message <- cli.Name + "login out"
			return
		case <-hasdata:
		case <-time.After((60 * time.Second)):
			delete(onlinemap, CliAddr)
			message <- cli.Name + "time out"
			return
		}
	}
}

// 等待总管道有数据写入
func Mannger() {
	onlinemap = make(map[string]Client)
	for {
		msg := <-message
		for _, val := range onlinemap {
			val.C <- msg
		}

	}
}

func main() {
	//1.监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listener.Close()

	//新开一个协程，转发消息，给每个成员发送消息
	go Mannger()

	//2主协程，循环阻塞等待用户链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		go HandleConn(conn)
	}

}
