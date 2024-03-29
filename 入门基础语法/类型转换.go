package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag=%t\n", flag)
	//bool类型不能转化为int!整型也不能转化为bool，不兼容
	// fmt.Printf("flag=%d\n", int(flag))

	//转化是显式的，不支持隐式
	var ch byte
	ch = 'a'
	var t int
	t = int(ch)
	fmt.Println("t= ", t)

}
