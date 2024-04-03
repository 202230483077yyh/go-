package main

import "fmt"

func main() {
	//有多少个[]就是多少维
	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]= %d	", i, j, a[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Println(a)

	//部分初始化，没有初始化的为0，记住后面还有，
	b := [3][4]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3},
	}
	fmt.Println(b)
}
