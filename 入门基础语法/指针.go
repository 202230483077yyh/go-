package main

import "fmt"

func main() {
	var p *int
	var i int
	i = 10
	p = &i

	fmt.Printf("i=  %d \n", *p)
	*p = 89
	fmt.Printf("i=  %d\n", i)
	fmt.Printf("p= %v\n", p)

	var tmp *int
	fmt.Println("tmp=", tmp)
	//等价于tmp=nil

}
