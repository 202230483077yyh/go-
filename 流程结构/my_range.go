package main

import "fmt"

func main() {
	str := "abcde"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Printf("\n")

	//range返回2个参数下标，元素值，也可以用匿名丢弃
	for i, c := range str {
		fmt.Printf("str[%d] = %c\n", i, c)
	}

}
