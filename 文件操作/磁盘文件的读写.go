package main

//可以拷贝视频、图片文件
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		_, err1 := f.WriteString(buf)
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			return
		}
	}
}

func ReadFile(path string) {
	//打开文件
	//f, err := os.Create(path)		Create会清空文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//推迟关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2)
	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错并且同时没到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))

}

func ReadFileline(path string) {
	//打开文件
	//f, err := os.Create(path)		Create会清空文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//推迟关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	//遇到‘\n’，结束读取,但是'\n'也读取进来

	for {
		buf, err1 := r.ReadBytes('\n')
		if err1 != nil {
			if err1 == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err1 = ", err1)
		}
		fmt.Printf("buf = #%s", string(buf))
	}

}

func main() {
	path := "./emmo.txt"
	//WriteFile(path)

	//ReadFile(path)
	ReadFileline(path)
}
