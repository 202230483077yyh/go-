package main

import (
	"fmt"
	"math"
)

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag=%t\n", flag)
	//bool类型不能转化为int!整型也不能转化为bool，不兼容
	// fmt.Printf("flag=%d\n", int(flag))

	//go转化是显式的，不支持隐式
	var ch byte
	ch = 'a'
	var t int
	t = int(ch)
	fmt.Println("t= \n", t)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("f= %.2f\n", f)
}
