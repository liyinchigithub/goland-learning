package GoPointer

import (
	// "fmt"
	"log"
)

func GoPointerA()  {
	/*
	[指针地址]、[指针类型]
		一个指针变量可以指向任何一个值的内存地址，它所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节，占用字节的大小与所指向的值的大小无关。
		当一个指针被定义后没有分配到任何变量时，它的默认值为 nil。
		指针变量通常缩写为 [ptr]。
		每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
		Go语言中使用在变量名前面添加[&]操作符（前缀）来获取变量的内存地址（取地址操作）
	*/
	var v string
	ptr := &v    // v 的类型为 T
	// 其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做T的指针类型，*代表指针。
	log.Println(ptr)
	// 指针地址，[常用于]作为引用参数传递给函数 定义函数形参为指针类型即*string 调用函数作为入参时传递指针类型即&v
}

func GoPointerB()  {
	var cat int = 1
    var str string = "banana"
    log.Printf("%p %p\n", &cat, &str) // 使用 fmt.Printf 的动词[%p]打印 cat 和 str 变量的[内存地址]，指针的值是带有0x十六进制前缀的一组数据。
	// 输出0xc0000161a8 0xc000054530
}

func GoPointerC()  {
	// 准备一个字符串类型
    var house = "Malibu Point 10880, 90265"
    // 对字符串取地址, ptr类型为*string
    ptr := &house
    // 打印ptr的类型
    log.Printf("ptr type: %T\n", ptr)
    // 打印ptr的指针地址，使用%p
    log.Printf("address: %p\n", ptr)
    // 对指针进行取值操作
    value := *ptr
    // 取值后的类型
    log.Printf("value type: %T\n", value)
    // 指针取值后就是指向变量的值
    log.Printf("value: %s\n", value)
}

// 入参是变量的指针类型
func Exchangge(a,b *int)  {
	t:=*a
	*a=*b
	*b=t
}
	
// 指针变量交换值
func GoPointerExchange()  {
	var a,b =1,2
	Exchangge(&a,&b)
	log.Printf("a=%d,b=%d\n",a,b)
	log.Printf("a=%p,b=%p\n",&a,&b)

}