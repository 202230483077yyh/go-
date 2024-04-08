package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 1.23 7. 8.99 ls 66.6"

	//解释正则表达式	+匹配前一个字符的多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	//匹配多行div标签的字符串
	//reg :=regexp.MustCompile(`<div>(?s:(.*？))</div>`)
	if reg == nil {
		fmt.Println("regex error")
		return
	}

	//提取关键信息
	res := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("res = ", res)

	//过滤标签
	// for _, text := range res {
	// 	//fmt.Println("text[0] = ", text[0])//带标签
	// 	fmt.Println("text[1] = ", text[1]) //不代标签
	// }
}
