package main

import "fmt"

// 函数返回多个值
func myfunc1(x, y int) (jia, jian int) {
	jia = x + y
	jian = x - y
	return jia, jian
}

// 函数返回命名返回值
func myfunc2(x, y int) (t int) {
	t = x * y
	return
}

func main() {
	a, b := 20, 30
	c, d := myfunc1(a, b)
	fmt.Printf("c= %d, d= %d\n", c, d)
	e := myfunc2(a, b)
	fmt.Printf("e= %d\n", e)
}
