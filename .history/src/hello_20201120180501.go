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
       // 没有初始化就为零值
       var b int
       fmt.Println(b)
   
       // bool 零值为 false
       var c bool=true
       fmt.Println(c)
}