package GoStruct

import (
	"encoding/json"
	"log"
	"strconv"
)

/*
	[结构体]
		Go 语言中[数组]可以存储同一类型的数据，但不能存储不同的数据类型，这时需要使用[结构体]。
		结构体是由一系列具'有相同类型'或'不同类型'的数据构成的数据集合。
		结构体表示一项记录，比如：保存图书馆的书籍记录，每本书有以下属性：
			Title ：标题
			Author ： 作者
			Subject：学科
			ID：书籍ID
		结构体有属性也有方法，相当于Java类
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

		一旦定义了结构体类型，它就能用于变量的声明（实例化），语法格式如下：
			variable_name := structure_variable_type {value1, value2...valuen}
			或
			variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
		即将结构体赋值给某个变量
		Go中的struct结构类似于面向对象中的类。
		面向对象中，除了成员变量还有方法。
		Go中也有方法，它是一种特殊的函数，定义于struct之上(与struct关联、绑定)，被称为struct的receiver。
*/

type Books struct {
	title   string // 书的名称
	author  string // 书的作者
	subject string // 书的主题
	book_id int    // 书的id
}

type Student struct {
	name    string
	age     int
	address string
	sex     string
	isFat   bool
}

func GoStructA() {
	// 向结构体放入数据
	log.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 1})
	// 也可以使用 key => value 格式
	log.Println(Books{title: "Go 语言2", author: "www.runoob2.com", subject: "Go 语言教程2", book_id: 2})
	// 忽略的字段为 0 或 空（仅放入部分）
	log.Println(Books{title: "Go 语言3", author: "www.runoob3.com"})
	// 实例化结构体，并初始化所有字段
	stu := Student{"张三", 22, "xiamen", "boy", false}
	log.Println(stu)
}

/*
	[访问结构体成员]
	如果要访问结构体成员，需要使用点号 . 操作符，格式为：结构体.成员名
*/
func GoStructB() {
	//  声明结构体对象变量（实例化结构体）
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	// 初始化 Book1 结构体变量，写法相当于 var Book1 Books{title:"Go 语言",author:"www.runoob.com",ubject:"Go 语言教程",book_id:6495407}、var Book1 Books{"Go 语言","www.runoob.com","Go 语言教程",6495407}）
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	// 初始化 Book2 结构体变量
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	log.Printf("Book 1 title : %s\n", Book1.title)
	log.Printf("Book 1 author : %s\n", Book1.author)
	log.Printf("Book 1 subject : %s\n", Book1.subject)
	log.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	log.Printf("Book 2 title : %s\n", Book2.title)
	log.Printf("Book 2 author : %s\n", Book2.author)
	log.Printf("Book 2 subject : %s\n", Book2.subject)
	log.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

/*
	[结构体可以作为函数的参数]
	可以像其他数据类型一样，将结构体类型作为参数传递给函数。
	并以以上实例的方式访问结构体变量
*/

func GoStructC() {
	// 声明变量，变量类型是结构体
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

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
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
}

// 定义函数，入参是结构体类型
func printBook(book *Books) {
	log.Printf("Book title : %s\n", book.title)
	log.Printf("Book author : %s\n", book.author)
	log.Printf("Book subject : %s\n", book.subject)
	log.Printf("Book book_id : %d\n", book.book_id)
}

/*
	[结构体指针]
		可以定义指向结构体的指针类似于其他指针变量（指针变量就是使用内存地址存储的变量）
		格式如下：
			var struct_pointer *Books
			以上定义的指针变量可以存储结构体变量的地址。
				查看结构体变量地址，可以将 & 符号放置于结构体变量前：
			struct_pointer = &Book1
				使用结构体指针访问结构体成员，使用 "." 操作符：
			struct_pointer.title
*/
func GoStructD() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

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

func printBook2(book *Books) { // 获取变量指针值
	log.Printf("Book title : %s\n", book.title)
	log.Printf("Book author : %s\n", book.author)
	log.Printf("Book subject : %s\n", book.subject)
	log.Printf("Book book_id : %d\n", book.book_id)
}

/*
	结构体的方法（相当于Java类的方法）
*/

func (books *Books) GetBooks() string {
	log.Println("调用结构体方法")
	books.title = "未来世界是go的世界"
	books.author = "www.go.com"
	books.subject = "Go 语言教程"
	books.book_id = 6495407
	log.Printf("Book title : %s\n", books.title)
	log.Printf("Book author : %s\n", books.author)
	log.Printf("Book subject : %s\n", books.subject)
	log.Printf("Book book_id : %d\n", books.book_id)
	return `{"title":"` + books.title + `","author":"` + books.author + `","subject":"` + books.subject + `","book_id":"` + strconv.Itoa(books.book_id) + `"}`
}

//  调用结构体方法
func GoStructE() {
	// 定义一个新结构体
	type NewBooks struct {
		Title   string `json:"title"`
		Author  string `json:"author"`
		Subject string `json:"subject"`
		Book_id string `json:"book_id"`
	}
	// 实例化一个结构体对象
	var Book1 Books
	jsonStr := Book1.GetBooks() // 对象调用结构体方法
	log.Printf("调用结构体方法%s\n:", jsonStr)
	// 调用方法返回json字符串，将字符串转换成json对象
	var newBooks NewBooks                             // 实例化一个结构体对象
	err := json.Unmarshal([]byte(jsonStr), &newBooks) // 将json字符串转换成json对象
	if err != nil {
		log.Println("json转换失败")
		return
	}
	log.Printf("转换成json对象 %s\n", jsonStr)
}

/*
	结构体多种实例化方式：
	（1） 普通实例化方式

*/

// ① 格式：var 结构体实例化对象 结构体类型名=结构体类型名{属性列表}
func StructInitType01() {
	// 定义一个新结构体
	type Teacher struct {
		Name   string
		Age    int
		School string
	}

	// 实例化一个结构体对象
	var t Teacher = Teacher{}

	log.Printf("数据类型是：%T", t) // 数据类型是：GoStruct.Teacher  %T表示类型
	log.Println(t)            // 输出：{0}
	// 设置老师的名字、年龄、学校
	t.Name = "赵册珊"
	t.Age = 31
	t.School = "黑龙江大学"
	log.Printf("t为%v", t) // 输出：t {赵册珊 31 黑龙江大学} %v表示值（原值输出）
	// 实例化一个结构体对象
	var t1 Teacher = Teacher{"李哈哈", 22, "东北师范大学"}
	log.Printf("t1为%v", t1) // %v表示值（原值输出）
}

// ② 格式：var 实例化对象名 *结构体类型名 =new(结构体类型名)
func StructInitType02() {
	// 定义一个新结构体
	type Teacher struct {
		Name   string
		Age    int
		School string
	}

	// new方法实例化一个结构体对象
	var t *Teacher = new(Teacher)

	// t是指针，t其实指向的就是地址，应该给这个地址的指向的对象的字段威值
	(*t).Name = "李哈哈"
	(*t).Age = 22 //*的作用;根据地址取值
	// 为了符合程产员的编程习惯。go挺供了简化的赋值方式:
	// 原因:go的设计者为了程序员使用方便，底层会对 t.School = "清华大学" 进行处理 会给t加上取值运算（*t).School = "清华大学"
	t.School = "清华大学" //go编译器底层对t.School转化（*t).School = "清华大学"
	log.Println(*t)
}

// ③ 格式：var 实例化对象名 *结构体类型名 =&结构体类型名{属性列表}
func StructInitType03() {
	// 定义一个新结构体
	type Teacher struct {
		Name   string
		Age    int
		School string
	}

	// 创建老师结构体的实例、对象、变量:
	var t1 *Teacher = &Teacher{"马士:兵", 46, "洁华大学"}
	log.Println(*t1)

	var t *Teacher = &Teacher{}
	(*t).Name = "李擦擦"
	(*t).Age = 22
	t.School = "清华大学"
	log.Println(*t)
}
