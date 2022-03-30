package GoStructJSON

import (
	"fmt"
	"encoding/json"
	// "math/rand"
	
)

/*
	json.Marshal(struct变量)  结构体[转]json字符串（字符串转json）
	参考：https://blog.csdn.net/zxy_666/article/details/80173288
*/ 
func GoStruct2JSON()  {
	// 结构体
	type Student struct {
		Id   int	`json:"id"` // 如果变量打上了json标签，如：Name旁边的 `json:"name"` ，那么转化成的json key就用该标签“name”，否则取变量名作为key，如“Age”，“HIgh”。如果定义为小写字母开头 私有属性不能被json包访问
		Name string	`json:"name"`// [注意]如果定义为小写字母开头 私有属性不能被json包访问
		Age  int `json:"age"`	
		Sno  string	`json:"son"`
	}
	
	// 实例化一个数据结构，用于生成json字符串(相当于类的实例化)
	var s = Student{
		Id:   11,
		Name: "小王",
		Age:  10,
		Sno:  "n201205",
	}
	fmt.Printf("%#v\n", s) //	GoStructJSON.Student{Id:11, Name:"小王", Age:10, Sno:"n201205"}
	// struct转json字符串
	jsonByte,_ := json.Marshal(s) 
	//Marshal失败时err!=nil
    // if err != nil {
    //     fmt.Println("生成json字符串错误")
    // }
	// json转字符串
	jsonStr := string(jsonByte)// jsonByte是[]byte类型，转化成string类型便于查看
	fmt.Println(jsonStr) //	{"id":11,"name":"小王","age":10,"son":"n201205"}
	
	
}

/*
	json.Unmarshal()	json字符串[转]结构体
*/ 
func GoJSON2Struct()  {
	// 定义结构体
	type Teacher struct {
		Id   int
		Name string
		Age  int
		Tno  string
	}
	// json字符串
	var str = `{"Id":11,"Name":"王老师","Age":30,"Tno":"t201205"}`// 也可以写成"{\"Id\":11,\"Name\":\"王老师\",\"Age\":30,\"Tno\":\"t201205\"}"
	// 实例化结构体
	var tea Teacher
	// json字符串[转]结构体
	err := json.Unmarshal([]byte(str),&tea)
	// 
	if err !=nil{
		fmt.Println(err)
	}
	// 
	fmt.Printf("%#v\n",tea)// GoJSON2Struct.Teacher{Id:11, Name:"王老师", Age:30, Tno:"t201205"}
	// 获取结构体属性的值
	fmt.Println(tea.Tno)//t201205
	
}



/*
	json.Unmarshal()
	[嵌套结构]json字符串转结构体
	https://blog.csdn.net/wangyufeng43400141/article/details/107139182
*/ 


func GoJSON2SubNode()  {
	type Class struct {
		Name  string
		Grade int
	}

	type Stu struct {
		Name  string `json:"name"`
		Age   int
		HIgh  bool
		sex   string
		Class *Class `json:"class"`
	}
	

	//实例化一个数据结构，用于生成json字符串
    stu := Stu{
        Name: "张三",
        Age:  18,
        HIgh: true,
        sex:  "男",
    }

    //指针变量
    cla := new(Class) // 实例化指针结构体
    cla.Name = "1班"
    cla.Grade = 3
    stu.Class=cla

    //Marshal失败时err!=nil
    jsonStu, err := json.Marshal(stu)
	// 
    if err != nil {
        fmt.Println("生成json字符串错误")
    }

    //jsonStu是[]byte类型，转化成string类型便于查看
    fmt.Println(string(jsonStu))
	// 输出：{"name":"张三","Age":18,"HIgh":true,"class":{"Name":"1班","Grade":3}}
	/*
	* 只要是可导出成员（变量首字母大写），都可以转成json。因成员变量sex是不可导出的，故无法转成json。
	* 如果变量打上了json标签，如Name旁边的 `json:"name"` ，那么转化成的json key就用该标签“name”，否则取变量名作为key，如“Age”，“HIgh”。
	* bool类型也是可以直接转换为json的value值。Channel， complex 以及函数不能被编码json字符串。当然，循环的数据结构也不行，它会导致marshal陷入死循环。
	* 指针变量，编码时自动转换为它所指向的值，如cla变量。（当然，不传指针，Stu struct的成员Class如果换成Class struct类型，效果也是一模一样的。只不过指针更快，且能节省内存空间。）
	* ：json编码成字符串后就是纯粹的字符串了
	*/
}

/*
	interface{}类型其实是个空接口，即没有方法的接口。
	go的每一种类型都实现了该接口。
	因此，任何其他类型的数据都可以赋值给interface{}类型
*/ 
func GoInterface(){
	type Stu struct {
		Name  interface{} `json:"name"`
		Age   interface{}
		HIgh  interface{}
		sex   interface{}
		Class interface{} `json:"class"`
	}
	
	type Class struct {
		Name  string
		Grade int
	}

	//实例化一个数据结构，用于生成json字符串
    stu := Stu{
        Name: "张三",
        Age:  18,
        HIgh: true,
        sex:  "男",
    }

    //指针变量
    cla := new(Class)
    cla.Name = "1班"
    cla.Grade = 3
    stu.Class=cla

    //Marshal失败时err!=nil
    jsonStu, err := json.Marshal(stu)
    if err != nil {
        fmt.Println("生成json字符串错误")
    }

    //jsonStu是[]byte类型，转化成string类型便于查看
    fmt.Println(string(jsonStu)) //{"name":"张三","Age":18,"HIgh":true,"class":{"Name":"1班","Grade":3}}

}