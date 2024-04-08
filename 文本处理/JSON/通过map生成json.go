package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcase"
	m["subjects"] = []string{"go", "math"}
	m["price"] = 66.6
	m["isok"] = true
	//非格式化
	//result, err := json.Marshal(m)
	//格式化，第三个参数表示行左距
	result, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))

}
