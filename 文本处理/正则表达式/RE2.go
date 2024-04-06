package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 1.23 7. 8.99 ls 66.6"

	//解释正则表达式	+匹配前一个字符的多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("regex error")
		return
	}

	//提取关键信息
	res := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("res = ", res)
}
