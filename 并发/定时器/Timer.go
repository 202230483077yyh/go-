package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间为2s，2s后往time通道写内容
	timer := time.NewTimer(2 * time.Second) //这里的定时器只会响应一次
	fmt.Println("当前时间： ", time.Now())

	t := <-timer.C
	fmt.Println("t = ", t)

}
