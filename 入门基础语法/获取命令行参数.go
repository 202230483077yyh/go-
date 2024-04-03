package main

import (
	"fmt"
	"os"
)

func main() {
	//接受用户传递的参数，都是以字符串方式传递的
	list := os.Args
	n := len(list)
	fmt.Println("n=", n)
	//两种打印参数的方式
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}
	fmt.Printf("\n")
	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}

}
