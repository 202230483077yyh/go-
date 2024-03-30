package main

import "fmt"

func main() {
	p := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 24}
	h := p[2:7]

	for i := 0; i < 5; i++ {
		fmt.Println(h[i])
	}
	h[0] = 345
	fmt.Println(p)

	//结构体数组创建切片
	s := []struct {
		a int
		b bool
	}{
		{1, true}, {2, false}, {56, true}, //注意这里还有一个逗号
	}
	fmt.Println(s)

}
