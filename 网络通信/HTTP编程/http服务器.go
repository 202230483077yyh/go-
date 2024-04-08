package main

import (
	"fmt"
	"net/http"
)

// w给客户端发送数据
// req读取客户端发送的数据
func HandConn(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req.url = ", req.URL)
	fmt.Println("req.Header = ", req.Header)
	fmt.Println("req.Body = ", req.Body)
	fmt.Println("req.Method = ", req.Method)
	w.Write([]byte("hello go"))
}

func main() {
	//注册处理函数，只要用户链接，自动调用指定的处理函数
	http.HandleFunc("/mydream", HandConn)

	//监听绑定
	http.ListenAndServe(":8000", nil)
}
