package main

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
   // 字符串变量
   var a string = "Runoob"
   fmt.Println(a)
   // 整型变量
   var b, c int = 1, 2 // 声明2个变量并初始化值
   fmt.Println(b,c)
   var i int = 1
   var f float64 =1.1
   // var b bool = true
   var s string = "ha"
   fmt.Printf("%v %v %v %q\n", i, f, b, s)

   // 省略 var
   v_name := 1
   fmt.Println("v_name",v_name);
}