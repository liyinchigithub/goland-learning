package GoFor

import "fmt"

/* 
Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号
[init] 一般为赋值表达式，给控制变量赋初值； 三种写法：可省略、var i int=0 、 i:=0
[condition] 关系表达式或逻辑表达式，循环控制条件；
[post] 一般为赋值表达式，给控制变量增量或减量。可省略
*/

// 和 C 语言的 for 一样：for init; condition; post { }
func GoForA(){
		j:=10
		for i := 0; i < j; i++ {
			fmt.Printf("i为%d\n",i)
		}
		
}
// 和 C 的 while 一样：for condition { }   不写初始化和自增
func GoForB(){
		var i,j int=0,10
		for i<j{
			fmt.Printf("i为%d\n",i)
			i++
		}
}


// init 和 post 参数是可选的，我们可以直接省略它，类似 While 语句。
// 和 C 的 for(;;) 一样：for { }
func GoForC(){
	i:=0
	for {
		i++
		fmt.Printf("i为%d\n",i)
		if i==10 {
			fmt.Printf("i为%d，退出无限循环\n",i)
			break
		}
	}
}
// For-each range 循环 这种格式的循环可以对[字符串]、[数组]、[切片]等进行迭代输出元素
func GoForD(){
	arr:=[]string{"a","b","c"}
	for i,j:= range arr {
		fmt.Println(i,j) // i是索引下标，j表示每个数组的元素
	}
	/*
	输出
	0 a
	1 b
	2 c
	*/
}
// for-each range
func GoForE()  {
	number:=[6]int{1,2,3,4,5,6}
	for i, v := range number {
		fmt.Println(i,v)
		/*
		输出：
		0 1
		1 2
		2 3
		3 4
		4 5
		5 6
		*/
	}
}

// 循环嵌套 99乘法表
func GoForF(){
	for i := 1; i < 10; i++ {
		/*    fmt.Printf("第%d次：\n",m) */
			for j := 1; j <= i; j++ {
				fmt.Printf("%dx%d=%d ",i,j,i*j)
			}
			fmt.Println("")
		}
}



// 无线循环
func GoForG(){
	for true  {
        fmt.Printf("这是无限循环。\n");
    }
}

// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
func GoForH(){
	
}


/* 
循环控制语句可以控制循环体内语句的执行过程
[break] 经常用于中断当前 for 循环或跳出 switch 语句
[continue] 跳过当前循环的剩余语句，然后继续进行下一轮循环
[goto] 将控制转移到被标记的语句
*/