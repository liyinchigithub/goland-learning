package GoStrings

import (
	"fmt"
	"strings"
	"strconv"
)

/*
	Go常用的字符串处理函数 Strings Contains、Split、Join、Index、Replace、Trim、Fields、Repeat
	https://www.cnblogs.com/yang-2018/p/11112996.html
*/

//判断字符串s是否包含子串
func GoStringsContains() {
	var s string = "hello go"
	r := strings.Contains(s, "go") // true
	fmt.Println(r)                 //true
}

// strings.Split() 字符串分割成切片
func GoStringsSplit() {
	str1 := "123-234-54543" // 有符号规律的字符串
	arr := strings.Split(str1, "-")
	fmt.Println(arr) //[123 234 54543]
}

// strings.Join() 把切片连接成字符串
func GoStringsJoin() {
	str1 := "123-234-54543"
	arr := strings.Split(str1, "-") //[123 234 54543]
	str2 := strings.Join(arr, "@")  // 数组根据符号拼接成字符串
	fmt.Println(str2)               //123@234@54543
}

//将一系列字符串连接为一个字符串，之间用sep来分隔
func GoStringsJoinArray() {
	var s1 []string = []string{"1", "2", "3", "4"}
	s2 := strings.Join(s1, ",") //把切片连接成字符串
	fmt.Println(s2)             //1,2,3,4
}

//子串sep在字符串s中[第一次出现的位置]，存在返回0，不存在则返回-1
func GoStringsIndex() {
	var s string = "hello go"
	n1 := strings.Index(s, "go")
	fmt.Println(n1) //6
}

//返回count个s串联的字符串
func GoStringsRepeat() {
	var s string = "hello go"
	s3 := strings.Repeat(s, 2)
	fmt.Println(s3) //hello gohello go
}

//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
func GoStringsReplace() {
	var s string = "hello go"
	s4 := strings.Replace(s, "o", "e", -1)
	fmt.Println(s4) //helle ge

}

//字符串切割，如果后一个参数为空，则按字符切割，返回字符串切片
func GoStringsSplitB() {
	var s string = "hello go"
	s5 := strings.Split(s, " ")
	fmt.Println(s5) //[hello go]
}

//切除字符串两端的某类字符串
func GoStringsTrim() {
	var s string = "hello go"
	s6 := strings.Trim(s, "o")
	fmt.Println(s6) //hello g
}

//去除字符串的空格符，并且按照空格分割返回slice
func GoStringsFields() {
	s7 := " hello go "
	s8 := strings.Fields(s7)
	fmt.Println(s8) //[hello go]
}

/*
	[格式化字符串]
	Go 语言中使用 fmt.Sprintf 格式化字符串并赋值给新串
	https://www.runoob.com/go/go-basic-syntax.html
*/
func formatStringA() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"                    // %d 表示整型数字，%s 表示字符串
	var result = fmt.Sprintf(url, stockcode, enddate) // 输出：Code=123&endDate=2020-12-31
	fmt.Println(result)
}

func formatStringB() {
	// golang中用单引号 ’ 定义字符 类型属于 int	
	var  a = 'a'
	fmt.Printf("%v %T",a,a) //值是97 代表ASCII码 类型int32
	fmt.Printf("%c %T",a,a) //值是a 代表 原样    类型int32
	// 定义一个字符串输出里面的字符
	var b = "this"
	fmt.Printf("%v %c %T",b[2],b[2],b[2]) // 105 i uint8

	// 一个汉字占 3 个字节utf-8编码 一个字母占用 1 个字节 ASCII码
	// unsafe.Sizeof() 不能查看string类型的占用的存储空间 用len()
	var c = "this"
	fmt.Println(len(c)) // 占用4 个字节
	c = "你好hello"
	fmt.Println(len(c)) // 占用11个字节
	var d = '国'
	fmt.Printf("%v %c %T\n",d,d,d) // 22269(utf-8编码) 国 int32

}


func StringRune()  {
	var n = "你好好  我的world"
	fmt.Printf("%T\n",n)
	for _, v := range n {//rune
		fmt.Printf("%v(%c) ",v,v) //20320(你) 22909(好) 22909(好) 32( ) 32() 25105(我) 30340(的) 119(w) 111(o) 114(r) 108(l) 100(d)
	}
	

}

func StringByte()  {
	var s = "big"
	var byteStr = []byte(s)//byte类型
	byteStr[0] = 's'
	fmt.Println(string(byteStr)) //sig
	
    s = "你好的 world"
    var runeStr = []rune(s)//rune类型
    runeStr[0] = '我'
	fmt.Println(string(runeStr)) //我好的 world

}

// 使用fmt.Sprintf 其它类型 转 string 类型
func Sprintf2String()  {
	var n int = 30
	var m float64 = 12.34
	var b bool = true
	var c byte = 'a'

	str1 := fmt.Sprintf("%d", n)
	fmt.Printf("%v--%T\n", str1, str1) // 30--string

	str2 := fmt.Sprintf("%.2f", m)
	fmt.Printf("%v--%T\n", str2, str2) // 12.34--string

	str3 := fmt.Sprintf("%t", b)
	fmt.Printf("%v--%T\n", str3, str3) // true--string

	str4 := fmt.Sprintf("%c", c)
	fmt.Printf("%v--%T\n", str4, str4) // a--string

}

// 使用strconv 转换string类型
func Strconv()  {
	 //整型转换
	 var n1 int64 = 40
	 str_n := strconv.FormatInt(n1, 10)
	 fmt.Printf("%v--%T\n", str_n, str_n) //40--string
 
	 //浮点型转换
	 var f1 float64 = 40.3213213
	 str_f := strconv.FormatFloat(f1, 'f', 2, 32)
	 fmt.Printf("%v--%T\n", str_f, str_f) //40.32--string
 
	 //布尔型转换 无意义
	 var b1 bool = false
	 str_b := strconv.FormatBool(b1)
	 fmt.Printf("%v--%T\n", str_b, str_b) //false--string
 
	 //byte转换 无意义
	 var c1 byte = 'a'
	 str_c := strconv.FormatUint(uint64(c1), 10)
	 fmt.Printf("%v--%T\n", str_c, str_c) //97--string

	//string 转 int
	var m1 = "12312321"
	num, _ := strconv.ParseInt(m1, 10, 64) // 10 代表进制 64 代表位数
	fmt.Printf("%v--%T\n", num, num) // 12312321--int64
	
	//string 转 flaot
	var fl = "123.354543"
	flo_num,_ :=strconv.ParseFloat(fl,64) //64 代表 位数
	fmt.Printf("%v--%T\n", flo_num, flo_num) //123.354543--float64

}