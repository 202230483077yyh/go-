package main

import "fmt"

func main() {
	//字符类型：byte,rune
	//整型：int32,int64
	//浮点型：float32,float64(自动推导默认64)
	//字符串：string
	//复数类型：complex64,complex128
	//局部也有默认初始值！

	var a bool
	a = true
	fmt.Println("a= ", a)

	var f1 = 3.14
	fmt.Println("f1= ", f1)

	var c byte
	c = 't'
	fmt.Printf("c= %c\n", c)

	var str string
	str = "abc"
	fmt.Printf("str = %s\n", str)
	fmt.Printf("len: %d\n", len(str))

	//复数类型
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t= ", t)
	//分离实部虚部
	fmt.Println("real(t)=", real(t), "imag(t)=", imag(t))

}
