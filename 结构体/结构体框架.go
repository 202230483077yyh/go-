package main

import "fmt"

// 结构体就是由一系列字段组成，用.访问
// 另外结构体支持指针直接用.访问
type my_dream struct {
	a int
	b int
}

func main() {
	h := my_dream{23, 42}
	fmt.Printf("%d %d\n", h.a, h.b)
	p := &h
	p.a = 13
	fmt.Printf("%d\n", h.a)
}
