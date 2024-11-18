package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)

	//消费者，从channel读取内容打印
	go consumer(ch)

	//生产者，写入channel
	producer(ch) //channel传参，引用传递
}
