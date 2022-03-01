package GoStrings

import (
	"fmt"
	"strings"
)

/*
	Go常用的字符串处理函数 Strings Contains、Split、Join、Index、Replace、Trim、Fields
	https://www.cnblogs.com/yang-2018/p/11112996.html
*/

//判断字符串s是否包含子串
func GoStringsContains() {
	var s string = "hello go"
	r := strings.Contains(s, "go")// true
	fmt.Println(r) //true
}

// strings.Split() 字符串分割成切片
func GoStringsSplit()  {
	str1 := "123-234-54543"// 有符号规律的字符串
	arr := strings.Split(str1,"-")
	fmt.Println(arr) //[123 234 54543]
}

// strings.Join() 把切片连接成字符串
func GoStringsJoin()  {
	str1 := "123-234-54543"
	arr := strings.Split(str1,"-")//[123 234 54543]
	str2 := strings.Join(arr,"@")// 数组根据符号拼接成字符串
	fmt.Println(str2) //123@234@54543
}
	
//将一系列字符串连接为一个字符串，之间用sep来分隔
func GoStringsJoinArray()  {
	var s1 []string = []string{"1", "2", "3", "4"}
	s2 := strings.Join(s1, ",")//把切片连接成字符串
	fmt.Println(s2) //1,2,3,4
}
	

//子串sep在字符串s中[第一次出现的位置]，存在返回0，不存在则返回-1 
func GoStringsIndex()  {
		var s string = "hello go"
		n1 := strings.Index(s, "go")
		fmt.Println(n1) //6
}

//返回count个s串联的字符串
func GoStringsRepeat()  {
		var s string = "hello go"
		s3 := strings.Repeat(s, 2)
		fmt.Println(s3) //hello gohello go
}

//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
func GoStringsReplace()  {
	var s string = "hello go"
	s4 := strings.Replace(s, "o", "e", -1)
	fmt.Println(s4) //helle ge

}
	
//字符串切割，如果后一个参数为空，则按字符切割，返回字符串切片
func GoStringsSplitB()  {
	var s string = "hello go"
	s5 := strings.Split(s, " ")
	fmt.Println(s5) //[hello go]

}

//切除字符串两端的某类字符串
func GoStringsTrim()  {
	var s string = "hello go"
	s6 := strings.Trim(s, "o")
	fmt.Println(s6) //hello g
}

//去除字符串的空格符，并且按照空格分割返回slice
func GoStringsFields()  {
		s7 := " hello go "
		s8 := strings.Fields(s7)
		fmt.Println(s8) //[hello go]
}

/*
	[格式化字符串]
	Go 语言中使用 fmt.Sprintf 格式化字符串并赋值给新串
	https://www.runoob.com/go/go-basic-syntax.html
*/
func formatString() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"                    // %d 表示整型数字，%s 表示字符串
	var result = fmt.Sprintf(url, stockcode, enddate) // 输出：Code=123&endDate=2020-12-31
	fmt.Println(result)
}
