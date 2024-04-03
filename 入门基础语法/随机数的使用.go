package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子参数，只需一次
	//如果种子参数都一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子参数
	for i := 0; i < 5; i++ {
		//fmt.Println("rand= ", rand.Int())
		fmt.Println("rand= ", rand.Intn(100)) //随机生成限制在100内的数
	}

}
