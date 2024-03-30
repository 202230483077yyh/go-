package main

import "fmt"

func main() {
	//1.无初始化
	var a [2]string
	a[0] = "abcd"
	a[1] = "defr"

	for i := 0; i < 2; i++ {
		fmt.Printf("a[%d]: %s\n", i, a[i])
	}

	//2.有初始化
	p := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Printf("p[%d]: %d \n", i, p[i])
	}
	fmt.Println(a, p)
}
