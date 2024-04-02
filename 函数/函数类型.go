package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

// 函数也是一种数据类型，同故宫type给一个函数类型起名
// My_FuncType它是一个函数类型
type My_FuncType func(int, int) int

func main() {
	var res int
	res = Add(2, 3)
	fmt.Printf("res =%d\n", res)

	//声明一个函数类型的变量
	var test1 My_FuncType = Add
	res = test1(4, 5)
	fmt.Printf("res =%d\n", res)

	test1 = Minus
	res = test1(6, 2)
	fmt.Printf("res =%d\n", res)

}
