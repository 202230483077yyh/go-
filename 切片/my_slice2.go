package main

import "fmt"

func print_slice(s []int) {
	fmt.Println(s, len(s), cap(s))
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	//切片可以基于底层数组动态扩展，切片的切片可以把前一个切片当作数组
	s = s[:0]
	print_slice(s)

	s = s[:1]
	print_slice(s)

	s = s[:2]
	print_slice(s)

	s = s[1:4]
	print_slice(s)

}
