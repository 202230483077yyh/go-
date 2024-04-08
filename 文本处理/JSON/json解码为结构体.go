package main

import (
	"encoding/json"
	"fmt"
)

// 顺序无需对应，但是大小写敏感，该二次编码的编码
type IT struct {
	Company  string   `json:"company"` //二次编码
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

// 如果只想解析部分字段，可以在自定义结构体中保留所需成员变量
type IT2 struct {
	Company string `json:"company"` //二次编码
}

func main() {
	jsonbuf := `
	{
		"company": "itcase",
		"isok": true,
		"price": 66.6,
		"subjects": [
		"go",
		"math"
		]
	}
	`
	var tmp IT
	//记住转化为字符数组，而不是字符串！
	err := json.Unmarshal([]byte(jsonbuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp)

	var tmp2 IT2
	err2 := json.Unmarshal([]byte(jsonbuf), &tmp2)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("tmp2 = ", tmp2)

}
