package main

import "fmt"

type stu struct {
	name string
	id   int
}

// 空接口不包含任何方法，可以保存任意类型的值
// 所有类型都实现了空接口
func main() {
	var i interface{} = 1
	fmt.Println("i =", i)

	i = "abc"
	fmt.Println("i =", i)

	tmp := make([]interface{}, 3)
	tmp[0] = stu{"mike", 11}
	tmp[1] = 2
	tmp[2] = "yyh"

	for id, val := range tmp {
		//第一个返回的是值，第二个返回的是判断结果的真假
		if value, ok := val.(int); ok == true {
			fmt.Printf("tmp[%d] 类型为int ,内容为%d\n", id, value)
		} else if value, ok := val.(string); ok == true {
			fmt.Printf("tmp[%d] 类型为string ,内容为%s\n", id, value)
		} else if value, ok := val.(stu); ok == true {
			fmt.Printf("tmp[%d] 类型为stu ,内容为%v\n", id, value)
		}
	}

	for id, val := range tmp {
		switch value := val.(type) {
		case int:
			fmt.Printf("tmp[%d] 类型为int ,内容为%d\n", id, value)
		case string:
			fmt.Printf("tmp[%d] 类型为string ,内容为%s\n", id, value)
		case stu:
			fmt.Printf("tmp[%d] 类型为stu ,内容为%v\n", id, value)
		}
	}
}
