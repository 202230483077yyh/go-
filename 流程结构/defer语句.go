package main

import "fmt"

// defer只能用在函数或者方法内部
// defer推迟使用函数，但参数已经传入
// 推迟的函数先入栈，直到外层的函数全部执行完再执行（多个defer先进后出）
//注意哪怕延迟函数或者某个延迟调用发生错误，这些函数依旧会执行

func test(c int) int {
	res := 100 / c
	return res
}

func main() {
	//入门例子1
	// a := 10
	// defer fmt.Printf("a= %d ", a)
	// a++
	// defer fmt.Printf("a= %d ", a)
	// a++
	// fmt.Printf("a= %d ", a)

	//入门例子2
	// defer fmt.Println("AAAAAAAAAAAAA")
	// defer fmt.Println("BBBBBBBBBBBBB")
	// defer test(0)
	// defer fmt.Println("CCCCCCCCCCCCCC")

	//入门例子3
	a, b := 10, 20
	defer func(x int, y int) {
		fmt.Println("内部1	x=", x, "  y=", y)
	}(a, b)

	//优先看func后面的参数列表有无变量名
	defer func(a int, b int) {
		//这里的a,b 还是外部的a,b?（闭包特性）
		fmt.Println("内部2	a=", a, "  b=", b)
	}(a, b)

	defer func(int, int) {
		//这里的a,b 还是外部的a,b?（闭包特性）
		fmt.Println("内部3	a=", a, "  b=", b)
	}(a, b)
	a = 111
	b = 222
	fmt.Println("外部	a=", a, "  b=", b)

}
