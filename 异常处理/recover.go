package main

import "fmt"

// recover放在defer后面使用，恢复程序
func test1() {
	fmt.Println("aaaaaaaaa")
}

func test2(x int) {
	//设置recover
	defer func() {
		//recover()	//	不会导致错误信息
		//可以打印panic错误信息,如果有的话
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //别忘了()调用匿名函数

	var a [10]int
	a[x] = 111
}

func test3() {
	fmt.Println("cccccccccc")
}

func main() {
	test1()
	test2(9)
	test3()
}
