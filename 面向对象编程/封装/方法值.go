package main

import "fmt"

type person struct {
	age  int
	name string
}

func (p person) setvalue() {
	fmt.Println(p)
}

func (p *person) setpointer() {
	fmt.Println(*p)
}

func main() {
	r := person{11, "mike"}
	r.setpointer()
	//获取并保存方式入口地址--方法值，保存了接收者
	rFunc := r.setpointer
	rFunc()

	//方法表达式，没保存调用接收者，调用要传参
	rFunc2 := (person).setvalue
	rFunc2(r)
}
