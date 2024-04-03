package main

import "fmt"

// 数组做函数参数，它是值传递
// 实参数组的每个元素给形参数组拷贝一份
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify a= ", a)
}

// 数组指针做函数参数
func modify1(a *[5]int) {
	//数组指针解引用方式！
	(*a)[0] = 666
	fmt.Println("modify a= ", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify1(&a)

	fmt.Println("main a= ", a)

}
