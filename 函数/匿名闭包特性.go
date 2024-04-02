package main

import "fmt"

func main() {
	a := 10
	str := "abcd"

	//匿名闭包以引用方式捕获值
	func() {
		a = 100
		str = "opqrst"
		fmt.Printf("内部 a=%d str=%s\n", a, str)
	}()
	fmt.Printf("外部 a=%d str=%s\n", a, str)

}
