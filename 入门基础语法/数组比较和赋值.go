package main

import "fmt"

func main() {
	//数组支持比较，只支持==或者！=，比较是不是每个元素都一样
	//2个数组比较，数组类型要一样

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}

	fmt.Println("a==b", a == b)
	fmt.Println("a==c", a == c)

	//同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println(d)

}
