package GoPointer

import (
	"fmt"
)

func GoPointerA()  {
	/*
	[指针地址]和[指针类型]
	一个指针变量可以指向任何一个值的内存地址，它所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节，占用字节的大小与所指向的值的大小无关。
	当一个指针被定义后没有分配到任何变量时，它的默认值为 nil。
	指针变量通常缩写为 [ptr]。
	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
	Go语言中使用在变量名前面添加[&]操作符（前缀）来获取变量的内存地址（取地址操作）
	*/
	var v string
	ptr := &v    // v 的类型为 T
	// 其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做T的指针类型，*代表指针。
	fmt.Println(ptr)
}

func GoPointerB()  {
	var cat int = 1
    var str string = "banana"
    fmt.Printf("%p %p", &cat, &str) // 使用 fmt.Printf 的动词[%p]打印 cat 和 str 变量的[内存地址]，指针的值是带有0x十六进制前缀的一组数据。
	// 输出0xc0000161a8 0xc000054530
}

func GoPointerC()  {
	// 准备一个字符串类型
    var house = "Malibu Point 10880, 90265"
    // 对字符串取地址, ptr类型为*string
    ptr := &house
    // 打印ptr的类型
    fmt.Printf("ptr type: %T\n", ptr)
    // 打印ptr的指针地址
    fmt.Printf("address: %p\n", ptr)
    // 对指针进行取值操作
    value := *ptr
    // 取值后的类型
    fmt.Printf("value type: %T\n", value)
    // 指针取值后就是指向变量的值
    fmt.Printf("value: %s\n", value)
}