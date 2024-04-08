package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	id   int
	name string
}

func main() {
	var s student
	//同名时候，就近原则
	s.name = "mike"
	//显示调用赋值父类的同名字段
	s.person.name = "yyh"
	s.age = 10
	s.id = 2022
	fmt.Println("s = ", s)
}
