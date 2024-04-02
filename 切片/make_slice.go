package main

import "fmt"

func main() {
	//make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

	//make函数传入两个参数,指定数组类型，cap
	//make函数传入三个参数，指定数组类型，len ,cap

	a := make([]int, 0, 5)
	b := a[:5]
	print_Slice("a", a)
	print_Slice("b", b)
}

func print_Slice(s string, v []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(v), cap(v), v)
}
