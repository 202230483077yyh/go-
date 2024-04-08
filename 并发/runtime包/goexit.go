package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccc")
	//终止此函数
	//return
	//终止所在协程
	runtime.Goexit()
	fmt.Println("dddddddddddddd")
}

func main() {

	//创建新的协程
	go func() {
		fmt.Println("aaaaaaaaaaaa")
		//中间调用其他程序
		test()
		fmt.Println("bbbbbbbbbbbb")
	}()
	for {

	}
}
