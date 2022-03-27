package main // 包

import (
	"fmt"
	"os"
	"GoStruct" // 引入模块第一种方式：需要把包放入 /usr/local/go/src/GoStruct。
	"src/GoHttpClient"// 引入模块第二种方式：配合go.mod文件，需与main.go在同一个目录下。
) // 导入内置库
func main() {
	// 调试输出
	// fmt.Print("Hello World")
	// fmt.Println("hello world")
	// 字符串格式化

	// 数据类型

	// excel操作

	// 数据库操作

	// os
	var args = os.Args
	fmt.Println(args)
	if string(args[1]) == "ok" {
		fmt.Printf("命令行入参是：%s，是OK\n", args[1])
		// 第一种方式：将模块拷贝到目录/usr/local/go/src/GoStruct下，调用模块封装方法
		GoStruct.GoStructB()
		// 第二种方式：配置go.mod文件 module src 并引入模块 src/模块名
		GoHttpClient.GentlemanSampleRequest("http://httpbin.org/post") // Go gentleman
		os.Exit(0) //main方法退出
	} else {
		fmt.Println(args)
		fmt.Printf("命令行入参是：%s，不是OK\n", args[1])
		os.Exit(0)//main方法退出
	}

	
}


/*
	运行脚本：
	cd src
	go run main.go ok
*/


/*
package main 定义了包名，必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
func main() 是程序开始执行的函数。
main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
*/


