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
	// 结构体实例化
	var a =NokiaPhone{}
	a.call()

	phone = new(NokiaPhone)// 结构体实例化对象，接口类型变了可以存放实现该接口方法的实例对象
	phone.call()
	
	// 结构体实例化
	var b =IPhone{}
	b.call()

	phone = new(IPhone)// 结构体实例化对象，接口类型变了可以存放实现该接口方法的实例对象
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
	// 实例化结构体对象
	C:=A_struct{name:"李四",age:12}
	fmt.Println(C.Asay("王五",22))
	
	B:=B_struct{}
	B.AMove()

}


/*
	接口继承和重写
*/ 

type Employee struct {
    Name string
}

func (e Employee) Work() {
    println(e.Name, "work")
}

type Developer struct {
    Employee // 无名称字段类型实现继承
    Projects []string
}

// 实现重写
func (d Developer) Work() {
    println(d.Name, "work on projects", d.Projects)
}

type Worker interface {
    Work()
}

// 定义函数
func workerWork(worker Worker) {
    worker.Work()
}

func TestOveri(t *testing.T) {
    wongoo := Developer{
        Employee: Employee{Name: "wongoo"},
        Projects: []string{"gateway", "message"},// 切片
    }
    workerWork(wongoo)          // wongoo work on projects [gateway message]
    workerWork(wongoo.Employee) // wongoo work
}