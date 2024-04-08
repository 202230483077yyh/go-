package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage :xxx scrFile dstFile")
		return
	}
	srcFile := list[1]
	dstFile := list[2]
	if srcFile == dstFile {
		fmt.Println("源文件和目标文件不能相同")
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFile)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建的目标文件
	dF, err2 := os.Create(dstFile)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	defer sF.Close()
	defer dF.Close()

	//核心处理：从源文件读取内容，往目标文件写，读多少，写多少
	buf := make([]byte, 4*1024)
	for {
		n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		//往目标文件写，读多少写多少
		dF.Write(buf[:n])
	}

}
