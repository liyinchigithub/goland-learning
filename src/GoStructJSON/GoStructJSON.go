package GoStructJSON

import (
	"fmt"
	"encoding/json"
	"math/rand"
	
)

/*
	@method  json.Marshal(struct变量)
	@description 结构体转json字符串
*/ 
func GoStruct2JSON()  {
	type Student struct {
		Id   int
		Name string
		Age  int //如果定义为age 私有属性不能被json包访问
		Sno  string
	}
	
	//struct转json字符串
	var s = Student{
		Id:   11,
		Name: "小王",
		Age:  10,
		Sno:  "n201205",
	}
	fmt.Printf("%#v\n", s)//demo.Student{Id:11, Name:"小王", Age:10, Sno:"n201205"}
	jsonByte,_ := json.Marshal(s)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)//{"Id":11,"Name":"小王","Age":10,"Sno":"n201205"}
	
	
}

/*
@method   json.Unmarshal()
@description json字符串转结构体
*/ 
func GoJSON2Struct()  {
	type Teacher struct {
		Id   int
		Name string
		Age  int
		Tno  string
	}
	// json字符串
	var str = `{"Id":11,"Name":"王老师","Age":30,"Tno":"t201205"}`
	var tea Teacher
	err := json.Unmarshal([]byte(str),&tea)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Printf("%#v\n",tea)//demo.Teacher{Id:11, Name:"王老师", Age:30, Tno:"t201205"}
	fmt.Println(tea.Tno)//t201205
	
}



/*
@method   json.Unmarshal()
@description [嵌套结构]json字符串转结构体
@document https://blog.csdn.net/wangyufeng43400141/article/details/107139182
*/ 



/*
	@method   json.Unmarshal()
	@description [嵌套结构]json字符串转结构体
	@document https://blog.csdn.net/wangyufeng43400141/article/details/107139182
*/ 