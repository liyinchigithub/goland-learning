package GoArithmeticOperator

import "fmt"
// 算术运算符

// 除法
func chufa()  {
	var a = 10
	var b = 3
	fmt.Println(a / b) //3 都是int 返回int
	var c = 10.0
	fmt.Println(c / float64(b)) //3.3333333333333335 //都是float 返回float
}

// [取余] 余数 = 被除数 -（被除数/除数）*除数

func quyu()  {
	fmt.Println(10%3)//1
	fmt.Println(-10%3)//-1  -10 - （-10/3）*3 = -1
	fmt.Println(10%-3)//1   10 - (10/-3)*-3 = 1
}

// 只能单独使用 不能赋值 只有 n++ n-- 没有 ++n --n

func pp()  {
	var n = 10
	n++
	fmt.Println(n) // 11
	//var m = n error
}

// 位运算符 和 逻辑运算

func wei()  {
	var i = 5  // 二进制 101
	var y = 2  // 二进制 010
	fmt.Println(i&y)  //且  	0
	fmt.Println(i|y)  //或  	7
	fmt.Println(i^y)  //异或  	7
	fmt.Println(i<<y) //左移 i*2的y次  20
	fmt.Println(i>>y) //右移 i/2的y次  1

}