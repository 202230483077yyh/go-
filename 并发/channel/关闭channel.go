package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//如果不需要再写数据时关闭channel
		close(ch)
	}()

	//1.遍历管道内容1
	// for {
	// 	//ok获取管道是否关闭,就算关了有数据还是可以读的
	// 	if num, ok := <-ch; ok == true {
	// 		fmt.Println("num = ", num)
	// 	} else {
	// 		break
	// 	}
	// }

	//2.遍历管道内容2,关闭的时候自动跳出循环
	for num := range ch {
		fmt.Println(num)
	}

}
