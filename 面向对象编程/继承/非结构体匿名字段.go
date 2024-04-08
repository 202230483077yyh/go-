package main

import "fmt"

type mystr string

type person struct {
	name string
	age  int
}

type student struct {
	person
	int   //基础类型的匿名字段
	mystr //别名的匿名字段
}

// 当结构体包含结构体指针匿名字段
type student1 struct {
	*person
	int
	mystr
}

func main() {
	s := student{person{"mike", 16}, 2022, "hehe"}

	fmt.Printf("s = %+v\n", s)
	//直接调用类型名即可
	fmt.Println(s.person, s.int, s.mystr)

	//包含结构体匿名指针
	//1.直接初始化
	s1 := student1{&person{"john", 11}, 2023, "sb"}
	fmt.Println(s1.person.name, s1.person.age, s1.int, s1.mystr)

	//2.先定义变量
	var s2 student1
	s2.person = new(person)
	s2.name = "ttt"
	s2.age = 14
	s2.mystr = "xb"
	s2.int = 555
	fmt.Println("s2= ", s2)
}
