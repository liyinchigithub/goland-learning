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
// For-each range 循环 这种格式的循环可以对[字符串]、[数组]、[切片 slice]、[map]等进行迭代输出元素for 循环的 
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

// break 跳出整个循环
func GoForH(){
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
	   fmt.Printf("a 的值为 : %d\n", a);
	   a++;
	   if a > 15 {
		  /* 使用 break 语句跳出循环 */
		  break;
	   }
	}
}

// continue 仅跳过本次，执行下一个循环
func GoForI(){
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {// 省略初始化和自增，默认初始化a从0开始
	   if a == 15 {
		  /* 跳过此次循环 */
		  a = a + 1;
		  continue;
	   }
	   fmt.Printf("a 的值为 : %d\n", a);
	   a++;    
	}  
}

/*
goto 语句可以无条件地转移到过程中指定的行。
goto 语句通常与条件语句配合使用。
可用来实现[条件转移][构成循环][跳出循环体]等功能。
但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
*/ 

// 在变量 a 等于 15 的时候跳过本次循环并回到循环的开始语句 LOOP 处
func GoForJ(){
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
	LOOP: for a < 20 {
	   if a == 15 {
		  /* 跳过迭代 */
		  a = a + 1
		  goto LOOP
	   }
	   fmt.Printf("a的值为 : %d\n", a)
	   a++    
	}  
}


/* 
循环控制语句可以控制循环体内语句的执行过程
[break] 经常用于中断当前 for 循环或跳出 switch 语句
[continue] 跳过当前循环的剩余语句，然后继续进行下一轮循环
[goto] 将控制转移到被标记的语句
*/