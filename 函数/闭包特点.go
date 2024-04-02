package main

import "fmt"

// 函数的返回值是一个匿名函数，返回一个函数类型
func test2() func() int {
	var x int

	return func() int {
		x++ //变量一直都在
		return x * x
	}
}

func test1() int {
	//函数调用时x才分配空间初始化为0
	var x int
	x++
	return x * x //函数调用完毕，x自动释放
}

func main() {
	//fmt.Printf("%d\n", test1())
	//闭包不关心这些捕获的变量和常量是否超出了作用域
	//所以只要闭包还在使用它，这些变量就还在

	fmt.Println("调用test2")
	f := test2()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16

}
