package main

import (
	"fmt"
	"time"
)

//goroutine是协程，比线程更小

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}

// 注意：main主协程退出，其他子协程也跟着退出
func main() {
	//只需在函数调用前添加go，新建一个任务（协程）
	go newTask()
	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}

}
