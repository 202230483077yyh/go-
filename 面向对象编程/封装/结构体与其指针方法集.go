package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func (tmp person) setvalue(age int, name string) {
	tmp.age = age
	tmp.name = name
}

func (tmp *person) setpointer(age int, name string) {
	tmp.age = age
	tmp.name = name
}

func (tmp person) print() {
	fmt.Println(tmp)
}

func (tmp *person) print1() {
	fmt.Println(*tmp)
}

func main() {
	v := person{11, "mike"}
	h := &person{23, "nini"}

	//结构体与结构体指针在作为接收者时候可以互相转化，调用对方的方法
	v.setpointer(41, "laji")
	h.setvalue(66, "yyh")
	v.print1()
	h.print()

}
