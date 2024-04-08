package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	//重置时间
	timer.Reset(1 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印，定时器时间到")
	}()

	//停止定时器,也不会往channel写入数据
	//timer.Stop()

	for {

	}
}
