package main

import "fmt"

func test(tmp map[int]string) {
	delete(tmp, 3)
}

func main() {
	m := map[int]string{1: "mike", 2: "titi"}
	fmt.Println(m)
	test(m)
	fmt.Println(m)

}
