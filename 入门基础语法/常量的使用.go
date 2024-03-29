package main

import "fmt"

func main() {
	//变量声明var，常量声明const
	const a int = 10
	fmt.Println("a=", a)

	//自动推导常量,没有:
	const b = 10
	fmt.Println("b=", b)
	fmt.Printf("the type of b is %T\n", b)

	//多个不同常量类型(也可自动推导)
	const (
		i int     = 10
		j float64 = 3.14
	)
	fmt.Printf("i=%d,j=%f", i, j)

}
