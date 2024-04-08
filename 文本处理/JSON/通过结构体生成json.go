package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量首字母大写，才可以转json
// 如果希望最后json格式的变量小写，可加上`json:"company"`
type IT struct {
	Company  string   `json:"company"` //二次编码
	Subjects []string `json:"-"`       //此字段不会输出
	Isok     bool     `json:",string"` //最后将bool转化为string输出
	Price    float64
}

func main() {
	//初始化一个结构体变量
	s := IT{"itcase", []string{"go", "math"}, true, 66.6}
	//编码，根据内容生成json文本，返回字符数组[]byte,error
	//1.非格式化编码
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//打印字符数组，要转化为string
	fmt.Println("buf = ", string(buf))

	//2.格式化编码
	buf, err = json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//打印字符数组，要转化为string
	fmt.Println("buf = ", string(buf))

}
