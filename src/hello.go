package main

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("hello world")
   // 字符串变量
   var a string = "Runoob"
   fmt.Println(a)
   // 同时声明2个变量及类型，并初始化值
   var b, c int = 1, 2 // 
   fmt.Println(b,c)
   //  整形
   var i int = 1
   // 浮点型
   var f float64 =1.1
   // 布尔类型
   var d bool = true
   var s string = "ha"
   // 
   fmt.Printf("%v %v %v %q\n", i, f, d, s)

   // 声明变量，省略 var方式
   v_name := 1
   fmt.Println("v_name",v_name);
}