1.同一目录下
只能属于同一个包，包内其他文件定义的函数可以直接调用
1）命令行：    go build 目录路径        (使用go build，首先要有go.mod)
（也可以配置任务文件，快捷ctrl+shift+B）    都是生成可执行文件再./main
2）命令行：    go run main.go x.go(同时运行编译多个文件)



2.不同目录：
不同文件夹，包名不一样
包名定义的函数只有首字母大写才可被外部访问（小写是私有函数）
运行：
在每个文件夹下创建.go之后，再go mod init 包名，创建go.mod
在go.mod中编辑：
replace 导入的包名 => 路径/包名         严格注意空格！
或者 go mod edit -replace example.com/greetings=../greetings

在命令行：
go mod tidy
或者?
go get 路径/包名

然后可以运行

**注意：如果想使用别的包的函数、结构体类型、结构体成员函数名、类型名、成员变量名，首字母必须大写！（可见性）**


3.当主模块有多个依赖项时，只需在主模块的go.mod添加replace即可
注意每个包都至少在同一级目录下有自己的go.mod
（A导入B，B导入C，D，那么在A的mod中必须导入C，D）

