package main

import "fmt"

func main() {
	//给int64取别名bigint
	type bigint int64

	var a bigint
	fmt.Printf("the type of a is %T\n", a)

	type (
		long int64
		char rune
	)
	var b long
	fmt.Printf("the type of b is %T\n", b)

}
