package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

// 值传递
func test1(a student) {
	a.id = 10
	fmt.Println("test1 a= ", a)
}

// 引用传递
func test2(a *student) {
	a.id = 15
	fmt.Println("test1 *a= ", *a)
}

func main() {
	s := student{1, "mike", 16}
	test1(s)
	fmt.Println("main s= ", s)

	test2(&s)
	fmt.Println("main s= ", s)
}
