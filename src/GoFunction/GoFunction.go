package GoFunction

import (
	"fmt"
)

func GoFunctionA()  {
	var a,b int=100,200;
	fmt.Printf("%d",max(a,b))
}

func GoFunctionB()  {
	fmt.Println(swap("hello","world"))
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int
	
	if (num1 > num2) {
		result = num1
		} else {
			result = num2
		}
		return result
	}
	
	/* 函数返回多个值 */
	func swap(a,b string)(string,string) {
		return a,a+b
	}
	
	
	/*
	Go 语言每个项目最少有个 main() 函数。
	你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。
	函数声明告诉了编译器函数的名称，返回类型，和参数。
	Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。
	[函数定义]
	Go 语言函数定义格式如下：
	func function_name( [parameter list] ) [return_types] {
		函数体
	}
	[函数定义解析]：
	func：函数由 func 开始声明
	function_name：函数名称，参数列表和返回值类型构成了函数签名。
	parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
	return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
	函数体：函数定义的代码集合。
	
	[函数参数]
	函数如果使用参数，该变量可称为函数的形参。
	(1)形参：在定义函数时，定义在函数体内的局部变量。
	(2)实参：调用函数，可以通过两种方式来传递参数：
	(3)值传递：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	(4)引用传递：* 和  & 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	[函数用法]
	函数作为另外一个函数的实参	函数定义后可作为另外一个函数的实参数传入
	闭包	闭包是[匿名函数]，可在动态编程中使用
	方法	方法就是一个包含了接受者的函数
	*/
	
	/*
	值传递：是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	*/ 
	func GoFunctionC()  {
		/* 定义局部变量 */
		var a int = 100
		var b int = 200
		
		fmt.Printf("交换前 a 的值为 : %d\n", a ) // 100
		fmt.Printf("交换前 b 的值为 : %d\n", b ) // 200
		
		/* 通过调用函数来交换值 */
		swap2(a, b)
		
		fmt.Printf("交换后 a 的值 : %d\n", a ) // 100
		fmt.Printf("交换后 b 的值 : %d\n", b ) // 200
	}
	
	
	/* 定义相互交换值的函数 */
	func swap2(x, y int) int {
		// 声明一个变量
		var temp int
		
		temp = x /* 保存 x 的值 */
		x = y    /* 将 y 值赋给 x */
		y = temp /* 将 temp 值赋给 y*/
		
		return temp;
	}
	
	/*
	程序中使用的是值传递, 所以两个值并没有实现交互，我们可以使用 [引用传递] 来实现交换效果。
	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：
	*/
	
	func GoFunctionD()  {
		/* 定义局部变量 */
		var a int = 100
		var b int= 200
	 
		fmt.Printf("交换前，a 的值 : %d\n", a )
		fmt.Printf("交换前，b 的值 : %d\n", b )
	 
		/* 调用 swap() 函数
		* &a 指向 a 指针，a 变量的地址
		* &b 指向 b 指针，b 变量的地址
		*/
		swap3(&a, &b)// 当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是[指针取值]  swap3内部改变了a和b的值，将影响到外部的a和b
	 
		fmt.Printf("交换后，a 的值 : %d\n", a )
		fmt.Printf("交换后，b 的值 : %d\n", b )
	}
	
	func swap3(x *int, y *int) {// 入参是指针 也就是引用传递
	// 声明变量
	var temp int
	temp = *x    /* 保存 x 地址上的值 */
	*x = *y      /* 将 y 值赋给 x */
	*y = temp    /* 将 temp 值赋给 y */
 	}


/*
	 匿名函数
	 函数作为实参 Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。类似JavaScript匿名函数，将匿名函数赋值给一个变量。
*/ 
func GoFunctionE()  {
	// 定义时被调用
	a:=func(x int)int{
		return x+1
	}
	//调用函数A，函数A入参是函数B，函数B有返回值
	x(a(10))
}

func x(x int)  {
	fmt.Println(x)
}

/*
	[闭包]
	Go 语言支持匿名函数，可作为闭包。
	匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
	以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。
	该函数的目的是在闭包中递增 i 变量，代码如下：
*/ 
func GoFunctionF()  {
	 /* nextNumber 为一个函数，函数 i 为 0 */
	 nextNumber := getSequence()  

	 /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	 fmt.Println(nextNumber())// 输出 1
	 fmt.Println(nextNumber())// 输出 2
	 fmt.Println(nextNumber())// 输出 3
	 
	 /* 创建新的函数 nextNumber1，并查看结果 */
	 nextNumber1 := getSequence()  
	 fmt.Println(nextNumber1())// 输出 1
	 fmt.Println(nextNumber1())// 输出 2
}

// 匿名函数
func getSequence() func() int {
	i:=0
	return func() int {
	   i+=1
	  return i  
	}
 }

/*
	[函数方法]
	Go语言中同时有[函数]和[方法]。
	一个方法就是一个包含了接受者的函数，接受者可以是'命名类型'或者'结构体类型'的一个值或者是一个指针。
	所有给定类型的方法属于该类型的方法集。
	[理解]
	Go中的struct结构类似于面向对象中的类，面向对象中，除了成员变量还有方法。
	Go中也有方法，它是一种特殊的函数，定义于struct之上(与struct关联、绑定)，被称为struct的receiver。
	https://www.cnblogs.com/f-ck-need-u/p/9890624.html
*/ 

/* 定义结构体 */
type Circle struct {
	// 结构体内定义变量和变量类型（属性、属性类型）
	radius float64
  }

//该 getArea方法，属于Circle结构体实例化对象中的方法（类似Java class中 定义一个方法）
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
  }

func GoFunctionG()  {
	// 实例化结构体对象
	var c1 Circle	// c1 为 Circle 类型对象，具备 Circle 类型对象的所有方法
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())//圆的面积 =  314
}

