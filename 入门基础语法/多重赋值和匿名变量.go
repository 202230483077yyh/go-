package main

import "fmt"

// 函数可以返回多个返回值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	//交换2个变量
	a, b := 10, 20
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Println("a=", a, "b=", b)

	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i= %d,j=%d\n", i, j)

	//匿名变量_，丢弃数据不处理
	//匿名变量配合函数返回值使用，才有优势
	tmp, _ = i, j
	fmt.Printf("tmp = %d\n", tmp)

	var c, d, e int
	c, d, e = test()
	fmt.Println("c= ", c, "d=", d, "e=", e)

	d, _, _ = test()
	fmt.Println("d=", d)

}
