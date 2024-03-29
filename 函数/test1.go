package main

import "fmt"

func MY_func1(a int, b int) {
	fmt.Printf("a = %d ,b = %d\n", a, b)
}

func main() {
	a, b := 10, 20
	MY_func1(a, b)
}
