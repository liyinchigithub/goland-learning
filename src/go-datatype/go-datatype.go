package test_datatype

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
	例如：声明两个变量都是指针
	var a, b *int
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
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	// 省略式声明变量
	r := 1
	fmt.Println(r)
	// 声明变量，省略 var方式
	v_name := 1
	fmt.Println("v_name", v_name)
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
	//
	fmt.Println("")
}
