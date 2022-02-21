package GoStruct

import (
	"fmt"
)

/*
	[结构体]
	Go 语言中[数组]可以存储同一类型的数据，存储不同的数据类型使用[结构体]。
	结构体是由一系列具'有相同类型'或'不同类型'的数据构成的数据集合。
	结构体表示一项记录，比如：保存图书馆的书籍记录，每本书有以下属性：
		Title ：标题
		Author ： 作者
		Subject：学科
		ID：书籍ID
	[结构体定义]
	需要使用 type 和 struct 语句。
	type 语句设定了结构体的名称。
	struct 语句定义一个新的数据类型，结构体中有一个或多个成员。
			格式如下：
					type struct_variable_type struct {
					member definition
					member definition
					...
					member definition
					}
	一旦定义了结构体类型，它就[能用于变量的声明]，语法格式如下：
	variable_name := structure_variable_type {value1, value2...valuen}
	或
	variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
	即将结构体赋值给某个变量

	Go中的struct结构类似于面向对象中的类。
	面向对象中，除了成员变量还有方法。
	Go中也有方法，它是一种特殊的函数，定义于struct之上(与struct关联、绑定)，被称为struct的receiver。
*/

type Books struct {
	title string // 书的名称
	author string // 书的作者
	subject string // 书的主题
	book_id int		 // 书的id
 }

 type Student struct{
	 name string
	 age int
	 address string
	 sex string
	 isFat bool
 }
func GoStructA()  {
	// 创建一个新的结构体（向结构体放入数据）
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
    // 忽略的字段为 0 或 空（仅放入部分）
   	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	//
	stu:=Student{"张三",22,"xiamen","boy",false}
	fmt.Println(stu)
}

/*
	[访问 结构体成员]
	如果要访问结构体成员，需要使用点号 . 操作符，格式为：结构体.成员名
*/ 
func GoStructB()  {
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

/*
	[结构体作为函数参数]
	可以像其他数据类型一样将结构体类型作为参数传递给函数。
	并以以上实例的方式访问结构体变量
*/ 

func GoStructC()  {
	// 声明变量，变量类型是结构体
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   printBook(Book1)

   /* 打印 Book2 信息 */
   printBook(Book2)
}

func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
 }

/*
	[结构体指针]
	可以定义指向结构体的指针类似于其他指针变量，格式如下：
	var struct_pointer *Books
	以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
	struct_pointer = &Book1
	使用结构体指针访问结构体成员，使用 "." 操作符：
	struct_pointer.title
*/
func GoStructD()  {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   printBook2(&Book1) // 使用变量指针地址

   /* 打印 Book2 信息 */
   printBook2(&Book2) // 使用变量指针地址

   /* 
   输出：
	Book title : Go 语言
	Book author : www.runoob.com
	Book subject : Go 语言教程
	Book book_id : 6495407
	Book title : Python 教程
	Book author : www.runoob.com
	Book subject : Python 语言教程
	Book book_id : 6495700
   */
}

func printBook2( book *Books ) { // 获取变量指针值
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
 }


