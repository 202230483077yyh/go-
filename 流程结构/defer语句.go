package main

import "fmt"

// defer推迟使用函数，但参数已经传入
// 推迟的函数先入栈，直到外层的函数全部执行完再执行
func main() {
	a := 10
	defer fmt.Printf("a= %d ", a)
	a++
	defer fmt.Printf("a= %d ", a)
	a++
	fmt.Printf("a= %d ", a)
}
