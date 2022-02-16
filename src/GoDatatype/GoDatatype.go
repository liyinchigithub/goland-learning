package GoDatatype

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

func GoDatatype() {// 变量名必须大写，才能被同包下其他函数调用
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
	var a1, b1, c1 = 1, 2, 3
	fmt.Printf("a=%d,b=%d,c=%d\n", a1, b1, c1) // d十进制 b二进制 o八进制 x十六进制
	// [省略式]
	r := 1
	fmt.Println(r)
	// [省略var]
	v_name := 1
	fmt.Println("v_name", v_name)
	// [省略数据类型]
	var address="quanzhou"
	fmt.Println(address)
	// 数字
	var d1 int = 1
	fmt.Println(d1)
	// 字符串
	var address2 string = "xiamen"
	fmt.Println(address2)
	// 多行字符串
	var moreString=
	`
	第一行
	第二行
	第三行
	`
	fmt.Printf("多行字符串%v",moreString) // Printf格式化输出，%d数值、%v原样输出
	// 数字
	var g int = 1
	// 浮点型
	var h1 float64 = 1.1
	const h2 = .71828 // 0.71828 只写小数部分
	const h3 = 1.     // 1	只写整数部分
	// 布尔类型
	var i bool = true
	var j string = "ha"
	fmt.Printf("%v %v %v %q\n", g, h1, i, j)
	// // 如果需要经常做类似的转换，可以将转换的代码封装成一个函数,例如下面布尔类型转数字，如果b为真，btoi返回1；如果为假，btoi返回0；
	// func btoi(b bool) int {
	// 	if b {
	// 		return 1 // 注意：区别于Java、JavaScript 布尔类型和1、0不等价，需要进行封装返回映射
	// 	}
	// 	return 0
	// }
	// // 数字到布尔型的逆转
	// btoi(true)
	// // itob报告是否为非零。
	// func intTObool(i int) bool { 
	// 	return i != 0 
	// }
	// intTObool(1)
	// Go语言中[不允许将整型强制转换为布尔型]布尔型无法参与数值运算，也无法与其他类型进行转换。
	/*
		var n bool
		fmt.Println(int(n) * 2)// 报错”cannot convert n (type bool) to type int“
	*/
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
	var a11 *int
	var a22 []int
	var a33 map[string]int
	var a44 chan int
	// var a55 func(string) int
	var a66 error    // error 是接口
	fmt.Println(a11) // <nil>
	fmt.Println(a22) // []
	fmt.Println(a33) // map[]
	fmt.Println(a44) // <nil>
	// fmt.Println(a55) // <nil>
	fmt.Println(a66) // <nil>
}
