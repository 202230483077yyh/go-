package main

import "fmt"

func test(a int) {
	if a == 0 {
		return
	}
	test(a - 1)
	fmt.Printf("a= %d \n", a)
}

func add(a int) int {
	if a == 0 {
		return 0
	}

	return add(a-1) + a
}

func main() {
	b := 10
	test(b)
	c := 100
	c = add(c)
	fmt.Printf("c= %d \n", c)

}
