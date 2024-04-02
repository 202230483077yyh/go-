package main

import "fmt"

type My_func func(int, int) int

// 回调函数，函数有一个参数是函数类型，这个函数是回调函数
func Calc(a, b int, test1 My_func) int {
	return test1(a, b)
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func cheng(a, b int) int {
	return a * b
}

func main() {
	fmt.Printf("2+3=%d\n", Calc(2, 3, Add))
	fmt.Printf("7-4=%d\n", Calc(7, 4, Minus))
	fmt.Printf("2*4=%d\n", Calc(2, 4, cheng))

}
