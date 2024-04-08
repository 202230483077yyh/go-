package main

import (
	"encoding/json"
	"fmt"
)

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
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonbuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("m = ", m)

	//从万能接口到具体类型不能强制转化，只可以类型断言
	// var str string
	// str = string(m["company"])
	// fmt.Println(str)

	var str string

	for key, value := range m {
		//下面的data才可以用于直接赋值给具体类型
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("m[%s]的类型为string-->value = %s\n", key, str)
		case bool:
			fmt.Printf("m[%s]的类型为bool-->value = %v\n", key, data)
		case []string:
			fmt.Printf("m[%s]的类型为[]string-->value = %v\n", key, data)
		case float64:
			fmt.Printf("m[%s]的类型为float64-->value = %v\n", key, data)
		}

	}

}
