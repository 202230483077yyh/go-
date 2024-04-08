package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person //只有类型没有名字，匿名字段，继承person成员
	id     int
}

func main() {
	var s1 student = student{person{"mike", 16}, 2022}
	fmt.Println("s1 = ", s1)
	//%+v,显示更加详细
	fmt.Printf("s1 = %+v\n", s1)

	//指定成员初始化，没指定的默认
	s3 := student{id: 3033}
	fmt.Printf("s3 = %+v\n", s3)

	//成员操作
	s1.name = "yyh"
	fmt.Println("s1 = ", s1)

	s1.person = person{"niuniu", 12}
	fmt.Println("s1 = ", s1)
}
