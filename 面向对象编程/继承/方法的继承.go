package main

import "fmt"

type person struct {
	age  int
	name string
}

func (tmp *person) Print() {
	fmt.Println(tmp.age, tmp.name)
}

type student struct {
	person
	id int
}

// 继承之后方法重写
func (tmp *student) Print() {
	fmt.Println(tmp.person, tmp.id)
}

func main() {
	s := student{person{11, "mike"}, 2022}
	//同名重写函数的调用，就近原则
	s.Print()
	//也可以显式调用父类
	s.person.Print()
}
