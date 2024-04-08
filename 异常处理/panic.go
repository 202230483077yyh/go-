package main

import "fmt"

// 只把panic函数作为报告致命错误(指针问题、数组越界)的一种方式
// 当panic异常发生时，程序会中断执行
// 直接调用内置的panic函数也会引发panic异常

func test1() {
	fmt.Println("aaaaaaaaa")
}

func test2(x int) {
	// panic("this is a panic test")
	//让程序自己调用
	var a [10]int
	a[x] = 111
}

func test3() {
	fmt.Println("cccccccccc")
}

func main() {
	test1()
	test2(23)
	test3()
}
