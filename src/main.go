package main // 包

import "fmt" // 导入内置库

var name string = "李四" //定义全局变量，数据类型为字符串，并赋值

/*
	[数据类型]
	在 Go 编程语言中，数据类型用于声明函数和变量。
	数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。
	Go 语言按类别有以下几种数据类型：
	https://www.runoob.com/go/go-data-types.html
*/

func dataType() {
	var a bool = false
	var b = "data type is string" // 字符串，不定义数据类型直接赋值s
	var c int                     // 不赋值默认为0
	var d string                  // 不赋值默认为字符串为 ""（空字符串）
	println(a)
	println(b)
	println(c)
	println(d)
	var a1 *int
	var a2 []int
	var a3 map[string]int
	var a4 chan int
	var a5 func(string) int
	var a6 error    // error 是接口
	fmt.Println(a1) // <nil>
	fmt.Println(a2) // []
	fmt.Println(a3) // map[]
	fmt.Println(a4) // <nil>
	fmt.Println(a5) // <nil>
	fmt.Println(a6) // <nil>
}

func main() {
	student("张三")
	/* 这是我的第一个简单的程序 */
	fmt.Print("Hello World")
	fmt.Println("hello world")

	// 字符串格式化
	formatString()
	// 数据类型
	dataType()
}

/*
第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。

下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

下一行  是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。

下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
使用 fmt.Print("hello, world\n") 可以得到相同的结果。
Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。

当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
*/
