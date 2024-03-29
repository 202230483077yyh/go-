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

}
