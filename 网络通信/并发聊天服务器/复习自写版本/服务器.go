package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Adrr string
}

// map、管道 只有make才有空间，不然连数都存不了
var onlinemap = make(map[string]Client)
var message = make(chan string)

func fenmessage(cli Client, conn net.Conn) {
	for {
		str := <-cli.C
		conn.Write([]byte(str))
	}

}

func readfenmsg(cli Client, conn net.Conn, end chan bool) {
	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("connread err = ", err)
			return
		}
		str := string(buf[:n])
		if str == "exit" {
			tmp := cli.Name + " logout!"
			delete(onlinemap, cli.Adrr)
			message <- tmp
			conn.Close()
			end <- true
			return
		} else if str == "who" {
			cli.C <- "user list:\n"
			for _, val := range onlinemap {
				cli.C <- val.Name + "\n"
			}

		} else if str[:6] == "rename" {
			cli.C <- "输入新名字："
			m, _ := conn.Read(buf)
			cli.Name = string(buf[:m])
			cli.C <- "改名成功"
		} else {
			message <- cli.Name + " : " + str
		}
	}

}

func HandleConne(conn net.Conn) {
	cli := Client{make(chan string), conn.RemoteAddr().String(), conn.RemoteAddr().String()}

	onlinemap[conn.RemoteAddr().String()] = cli
	end := make(chan bool)

	go fenmessage(cli, conn)
	message <- cli.Name + "加入聊天室"
	go readfenmsg(cli, conn, end)

	<-end

}

func allmanager() {
	for {
		str := <-message
		for _, val := range onlinemap {
			val.C <- str
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listener.Close()

	go allmanager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		go HandleConne(conn)
	}
}
