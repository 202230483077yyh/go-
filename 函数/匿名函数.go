package main

import "fmt"

func main() {
	a := 10
	str := "abcde"

	//匿名函数没有函数名字，但是函数已经定义只是没调用
	//f1内部可以捕获外部同作用域下数据
	f1 := func() {
		fmt.Printf("a= %d\n", a)
		fmt.Printf("str= %s", str)
	}

	f1()

	//给一个函数类型起别名
	type My_func func()
	var f2 My_func
	f2 = f1

	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Printf("a= %d\n", a)
		fmt.Printf("str= %s", str)
	}() //后面的()代表调用此匿名函数

}
