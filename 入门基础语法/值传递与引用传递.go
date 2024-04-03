package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Println("函数：a=", a, "b= ", b)
}

func swap1(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println("main：a=", a, "b= ", b)

	swap1(&a, &b)
	fmt.Println("main：a=", a, "b= ", b)
}
