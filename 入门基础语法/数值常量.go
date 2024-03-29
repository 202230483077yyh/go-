package main

import "fmt"

func main() {
	const (
		//数值常量是高精度的 值。
		big   = 1 << 100
		small = big >> 99
	)

	fmt.Printf("small = %d\n", small)

}
