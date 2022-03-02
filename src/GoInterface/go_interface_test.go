package GoInterface

import (
	"testing"
	"fmt"
)

// 定义接口
type Phone interface {
    // 接口函数
    call()
}

// 定义结构体
type NokiaPhone struct {
}
// 结构体实现接口的方法（相当于java类实现接口方法）
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

// 定义结构体
type IPhone struct {
}
// 结构体实现接口方法（相当于java类实现接口方法）
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}


func Test(t *testing.T)  {
	// 定义一个接口类型变量
	var phone Phone
	// 实例化
	phone = new(NokiaPhone)// 结构体实例化对象，接口类型变了可以存放实现该接口方法的实例对象
	phone.call()
	// 实例化
	phone = new(IPhone)// 结构体实例化对象，接口类型变了可以存放实现该接口方法的实例对象
	phone.call()
}