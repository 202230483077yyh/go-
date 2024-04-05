package main

import "fmt"

// 结构体就是由一系列字段组成，用.访问
// 另外结构体支持指针直接用.访问.		(*p1).a和p1.a完全等价
type my_dream struct {
	//无var关键字
	a int
	b int
}

func main() {
	h := my_dream{23, 42}
	fmt.Printf("%d %d\n", h.a, h.b)
	p := &h
	p.a = 13
	fmt.Printf("%d\n", h.a)

	//通过new来申请一个地址
	p1 := new(my_dream)
	p1.a = 23
	p1.b = 46
	fmt.Println("p1 = ", p1)
}
