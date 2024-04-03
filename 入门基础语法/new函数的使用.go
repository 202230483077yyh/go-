package main

import "fmt"

func main() {
	var p *int

	//分配内存，自动内存回收，无须手动释放
	p = new(int)
	*p = 666
	fmt.Println("*p=", *p)

}
