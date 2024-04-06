package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "abc avd avc abn atc"
	//1)解析规则,如果成功返回解释器
	//可以用反引号``(键1左边)，也可用""
	reg := regexp.MustCompile(`a.c`)
	if reg == nil {
		fmt.Println("regex err")
		return
	}

	//2）根据规则提取关键信息,-1表示所有的
	res := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("res = ", res)

}
