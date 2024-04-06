package main

import (
	"fmt"
	"time"
)

// 管道创建
var ch = make(chan int)

func Printer(str string) {
	for _, val := range str {
		fmt.Printf("%c", val)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
	ch <- 666 //给管道写数据
}

func person2() {
	<-ch //从管道取数据接受，如果为空阻塞
	Printer("world")
}

func main() {
	go person1()
	go person2()
	//runtime.Gosched()
	for {

	}

}
