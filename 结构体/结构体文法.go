package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  //全部初始化
	v2 = Vertex{X: 1}  // 指定部分初始化
	v3 = Vertex{}      //全部默认初始化
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	//结构体可以直接打印，呈现{}字段
	fmt.Println(v1, p, v2, v3)
}
