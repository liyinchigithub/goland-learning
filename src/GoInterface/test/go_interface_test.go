package GoInterface

import (
	"fmt"
	"log"
	"testing"
)

/*
	[接口]
	包含一组方法，但是不能包含其他的字段，因为接口是一种类型，它的字段只能是方法。
	且不含方法具体实现内容，所以接口可以被实现，也可以被继承。
*/

// 定义接口
type Phone interface {
	call() // 此方法还是构体的方法
	photo(text string)
}

// 定义函数
func photo(text string) {
	fmt.Printf("I am photo, I can %s  you!", text)
}

// 定义结构体 NokiaPhone
type NokiaPhone struct {
	name  string
	num   int
	Phone // 结构体实现接口（匿名字段）
}

// 定义结构体 NokiaPhone 方法，同时实现接口Phone的方法（相当于java类实现接口方法）
func (nokiaPhone *NokiaPhone) call() {
	log.Println("I am Nokia, I can call you!")
	// 实例化结构体对象.属性  设置值
	nokiaPhone.name = "nokia" //
	nokiaPhone.num = 123
	// 获取对象.属性值
	log.Printf("nokiaPhone.name:%s", nokiaPhone.name) // nokia
	log.Printf("nokiaPhone.name:%d", nokiaPhone.num)  // 123

}

// 定义结构体
type IPhone struct {
	Phone // 匿名字段，结构体实现接口
}

// 定义结构体方法，同时实现接口方法（相当于java类实现接口方法）
func (iPhone *IPhone) call() {
	log.Println("I am iPhone, I can call you!")
}

func Test(t *testing.T) {
	var phone Phone      // 定义一个接口类型变量
	var a = NokiaPhone{} // 实例化结构体
	a.call()

	phone = new(NokiaPhone) // 实例化结构体对象，等同于 nokiaPhone NokiaPhone
	phone.call()            // I am Nokia, I can call you!
	phone.photo("拍照")       //

	var b = IPhone{} // 实例化结构体
	b.call()         // I am iPhone, I can call you!

	phone = new(IPhone) // 结构体实例化对象，接口类型变量phone 可以存放实现该接口方法的实例对象
	phone.call()
	phone.photo("拍照") // I am Nokia, I can call you!
}

/*
	接口嵌套
*/

// 定义接口
type A_interfaceer interface {
	Asay(name string, age int) (string, int)
	AMove()
}

// 定义接口
type B_interfaceer interface {
	// 接口B嵌套接口A（接口继承接口）
	A_interfaceer
}

// 定义结构体
type A_struct struct {
	name string
	age  int
}

// 定义结构体
type B_struct struct {
}

//	结构体A_struct实现了接口A的方法Asay()
func (a A_struct) Asay(name string, age int) (string, int) {
	return name, age
}

func (b B_struct) AMove() {
	log.Println("结构体B_struct实现了接口A_interfaceer的方法AMove()")
}

// 单元测试
func TestInterface(t *testing.T) {

	C := A_struct{name: "李四", age: 12} // 实例化结构体
	log.Println(C.Asay("王五", 22))      // 结构体实现接口方法

	B := B_struct{} // 实例化结构体
	B.AMove()       // 结构体实现接口方法

}

/*
	接口实现
	结构体继承、重写
*/

// 定义结构体[Employee]
type Employee struct {
	Name string // 定义结构体字段
}

// 定义结构体[Employee]的方法
func (e Employee) Work() {
	log.Println(e.Name, "work()")
}

// 定义结构体[Developer]
type Developer struct {
	Employee //	结构体继承
	Projects []string
}

/*
	实现重写接口方法
*/

// 结构体[Developer]的方法，因结构体[Developer]继承了结构体[Employee]，固方法名一样时，结构体[Employee]的方法被继承者[Developer]重写）
func (d Developer) Work() {
	log.Println(d.Name, "work on projects", d.Projects)
}

//  定义接口[Worker]
type Worker interface {
	Work() // 接口中的方法
}

//	定义函数，入参是接口类型
func workerWork(w Worker) {
	w.Work()
}

/*
	[结构体实现接口方法]
*/
type struct_test struct {
	Worker // 结构体实现接口方法
}

func TestStructImplementingInterface(t *testing.T) {
	// 实例化结构体
	var s struct_test
	//
	s.Worker = Developer{
		Employee: Employee{
			Name: "John",
		},
		Projects: []string{"project1", "project2"},
	}
	// 调用结构体实现接口方法
	s.Worker.Work() // 结构体实例对象.接口.方法
	// 调用函数
	workerWork(s.Worker)
}

//  单元测试
// func TestOver(t *testing.T) {
// 	wongoo := Developer{
// 		Employee: Employee{Name: "wongoo"},
// 		Projects: []string{"gateway", "message"}, // 切片
// 	}
// 	workerWork(wongoo)          // wongoo work on projects [gateway message]
// 	workerWork(wongoo.Employee) // wongoo work
// }
