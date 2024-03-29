package GoJson

/*
	[读取json文件]
	首先要能读取文件内容，读取文件内容可以使用os.Open()打开一个文件，
	然后再使用ioutil.ReadAll()读取文件的内容。
	[注意]ioutil.ReadAll()读取到的是字节，也就是[]bytes数据类型。
	如果你需要把[字节]转换成[字符]可以在使用string()函数转成字符。
	如这个例子，读取user.json文件并打印出内容
*/


import (
	"fmt"
	"io/ioutil" // 读取json
	"encoding/json" // 写入json
	"os"
  )
  
 

// 读取json
func JsonReadFile()  {
	// 打开json文件
	jsonFile, err := os.Open("../config/config.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要记得关闭
	defer jsonFile.Close()
	// 读取文件（返回字节）
	byteValue, _ := ioutil.ReadAll(jsonFile) // _值为nil
	// 字节转字符串
	fmt.Println(string(byteValue))
}


// 写入json
func JsonWriteFile()  {
	// 成功读取到了JSON文件的内容，现在我们就要解析它，通常比较喜欢使用struct类型做转换，所以我们先定义一个struct

	// 定义结构体
	type PersonInfo struct {
		Name    string
		age     int
		Sex     bool
		Hobbies []string	// 数组
	}
	// 实例化结构体对象，并初始化
	personInfo := []PersonInfo{
		{"David", 30, true, []string{"跑步", "读书", "看电影"}},
		{"Lee", 27, false, []string{"工作", "读书", "看电影"}}}

    // 创建文件
    filePtr, err := os.Create("../util/person_info.json")
	// 
    if err != nil {
        fmt.Println("Create file failed", err.Error())
        return
    }
    defer filePtr.Close()

    // 创建Json编码器
    encoder := json.NewEncoder(filePtr)

    err = encoder.Encode(personInfo)
    if err != nil {
        fmt.Println("Encoder failed", err.Error())

    } else {
        fmt.Println("Encoder success")
    }

// 带JSON缩进格式写文件
//    data, err := json.MarshalIndent(personInfo, "", "  ")
//    if err != nil {
//     fmt.Println("Encoder failed", err.Error())
   
//    } else {
//     fmt.Println("Encoder success")
//    }
   
//    filePtr.Write(data)
}


