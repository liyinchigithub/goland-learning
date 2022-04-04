package GoInterface

import (
	"testing"
	"fmt"
)

/*
	
*/

// 定义接口
type Phone interface {
    // 接口中的函数
    call()
}

// 定义结构体
type NokiaPhone struct {
	name string
	num int
}
// 结构体实现接口的方法（相当于java类实现接口方法）
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
	nokiaPhone.name = "nokia"
	nokiaPhone.num = 123
    fmt.Printf("nokiaPhone.name:%s",nokiaPhone.name)
    fmt.Printf("nokiaPhone.name:%d",nokiaPhone.num)
	
	
}

// 定义结构体
type IPhone struct {
}
// 结构体实现接口方法（相当于java类实现接口方法）
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}


func Test(t *testing.T)  {
	var phone Phone// 定义一个接口类型变量
	var a =NokiaPhone{}// 实例化结构体
	a.call()// I am Nokia, I can call you!

	phone = new(NokiaPhone)// 实例化结构体
	phone.call()// I am Nokia, I can call you!
	
	var b =IPhone{}// 实例化结构体
	b.call()// I am iPhone, I can call you!

	phone = new(IPhone)// 结构体实例化对象，接口类型变量可以存放实现该接口方法的实例对象
	phone.call()
}

/*
	接口嵌套
*/ 

// 定义接口
type A_interfaceer interface{
	Asay(name string ,age int)( string, int) 
	AMove()
}

// 定义接口
type B_interfaceer interface{
	// 接口B嵌套接口A
	A_interfaceer
}
// 定义结构体
type A_struct struct{
	name string
	age int
}

// 定义结构体
type B_struct struct{
}

//	结构体A_struct实现了接口A的方法Asay()
func (a A_struct )Asay(name string ,age int)( string, int)  {
	return name,age
}

func (b B_struct)AMove()  {
	fmt.Println("结构体B_struct实现了接口A_interfaceer的方法AMove()")
}

// 单元测试
func TestInterface(t *testing.T)  {

	C:=A_struct{name:"李四",age:12}// 实例化结构体
	fmt.Println(C.Asay("王五",22))// 结构体实现接口方法
	
	B:=B_struct{}// 实例化结构体
	B.AMove()// 结构体实现接口方法

}


/*
	接口继承和重写
*/ 

// 定义结构体 
type Employee struct {
    Name string
}

// 定义结构体的方法
func (e Employee) Work() {
    println(e.Name, "work")
}

// 定义结构体 
type Developer struct {
    Employee // 无名称字段类型实现继承
    Projects []string
}

// 实现重写
func (d Developer) Work() {
    println(d.Name, "work on projects", d.Projects)
}

//  接口继承
type Worker interface {
    Work()
}

//	定义函数
func workerWork(worker Worker) {
    worker.Work()
}

//  单元测试
func TestOver(t *testing.T) {
    wongoo := Developer{
        Employee: Employee{Name: "wongoo"},
        Projects: []string{"gateway", "message"},// 切片
    }
    workerWork(wongoo)          // wongoo work on projects [gateway message]
    workerWork(wongoo.Employee) // wongoo work
}