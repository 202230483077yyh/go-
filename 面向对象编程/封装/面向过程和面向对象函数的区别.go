package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

// 给某个类型绑定一个函数
type long int

// tmp是接收者，带有接收者的函数称为方法
func (tmp long) Add2(other long) long {
	return tmp + other
}

func main() {
	var a long
	a = 2
	//调用方法：变量名.函数（参数）
	res := a.Add2(5)
	fmt.Println("res = ", res)
}
