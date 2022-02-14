package go_datatype

import (
	"fmt"
)

/*
	[数据类型]
	Go语言是静态类型语言，因此变量（variable）是有明确类型的，编译器也会检查变量类型的正确性。
	在数学概念中，变量表示没有固定值且可改变的数。但从计算机系统实现角度来看，变量是一段或多段用来存储数据的内存。
	[声明变量]
	var name type
	var 是声明变量的关键字，name 是变量名，type 是变量的类型,明变量时将变量的类型放在变量的名称之后。
	例如：声明两个变量都是指针	var a, b *int
	[对数据类型转换比较严格，不支持隐式数据转换，只能显示转换]
*/
func test_datatype() {
	// 常量
	const pi = 3.14
	fmt.Println(pi)
	// 变量
	var name string = "张三"
	// 定义多个变量
	var (
		age    int     = 22
		blance float32 = 1.2
	)
	fmt.Println(name, age, blance)
	// 定义多个变量
	var a, b, c = 1, 2, 3
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c) // d十进制 b二进制 o八进制 x十六进制
	// [省略式]
	r := 1
	fmt.Println(r)
	// [省略var]
	v_name := 1
	fmt.Println("v_name", v_name)
	// [省略数据类型]
	address="quanzhou"
	fmt.Println(address)
	// 数字
	var d int = 1
	fmt.Println(d)
	// 字符串
	var address string = "xiamen"
	fmt.Println(address)
	// 数字
	var g int = 1
	// 浮点型
	var h float64 = 1.1
	// 布
	var i bool = true
	var j string = "ha"
	fmt.Printf("%v %v %v %q\n", g, h, i, j)
	// 数组
	// 对象
	//
	// 默认值
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
