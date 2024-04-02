package main

import "fmt"

func main() {
	//for 初始语句;判断条件;条件变化{

	//}
	num := 0
	for i := 0; i <= 100; i++ {
		num += i
	}
	fmt.Printf("num = %d\n", num)

	//for可以不写初始化语句和条件改变语句,类似while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//for什么都不写就是无限循环
	tmp := 1
	for {
		if tmp >= 10 {
			break
		}
		tmp = tmp + tmp
	}
	fmt.Printf("tmp = %d\n", tmp)

}
