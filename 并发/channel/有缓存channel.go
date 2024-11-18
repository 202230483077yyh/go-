package main

import (
	"fmt"
	"time"
)

func main() {
	//创建有缓存的channel，只有channel满的时候才会阻塞
	ch := make(chan int, 3)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i <= 3; i++ {
			ch <- i
			fmt.Printf("子协程[%d]:len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i <= 3; i++ {
		<-ch
		fmt.Printf("主协程[%d]:len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))

	}

}
