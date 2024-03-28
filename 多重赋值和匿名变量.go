package main

import "fmt"

func main() {
	a, b := 10, 20
	var tmp int
	tmp = a
	a = b
	b = tmp

	fmt.Println("a=", a, "b=", b)

}
