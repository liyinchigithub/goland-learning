package GoCondition

import (
	"fmt"
)

/*
	switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 
*/

// if
func GoConditionA()  {
		/* 定义局部变量 */
		var a int = 10
		/* 使用 if 语句判断布尔表达式 */
		if a < 20 {
			/* 如果条件为 true 则执行以下语句 */
			fmt.Printf("a 小于 20\n" )
		}
		fmt.Printf("a 的值为 : %d\n", a)
		
}

// if else
func GoConditionB(){
		/* 局部变量定义 */
		var a int = 100;
		/* 判断布尔表达式 */
		if a < 20 {
			/* 如果条件为 true 则执行以下语句 */
			fmt.Printf("a 小于 20\n" );
		} else {
			/* 如果条件为 false 则执行以下语句 */
			fmt.Printf("a 不小于 20\n" );
		}
		fmt.Printf("a 的值为 : %d\n", a);
	 
}

// if嵌套
func GoConditionC(){
		/* 定义局部变量 */
		var a int = 100
		var b int = 200
	
		/* 判断条件 */
		if a == 100 {
			/* if 条件语句为 true 执行 */
			if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
			}
		}
		fmt.Printf("a 值为 : %d\n", a );
		fmt.Printf("b 值为 : %d\n", b );
}

// switch 写法一
func GoConditionD(){
	var mark int=3;
	switch mark{
	case 1:fmt.Printf("mark is%d",mark)
	case 2:fmt.Printf("mark is%d",mark)
	case 3:fmt.Printf("mark is%d",mark)
	default:
		fmt.Printf("mark is%d",mark)
	}
}

// switch 写法二
func GoConditionE(){
	var mark int=2;
	switch  {
	case mark==1:fmt.Println("mark is 1")
	case mark==2:fmt.Println("mark is 2")
	case mark==3:fmt.Println("mark is 3")
	default:
		fmt.Printf("mark is %d",mark)
	}
}

// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true
func GoConditionF(){
	switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")// 停止不再下去
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }

	/*
	2、case 条件语句为 true
	3、case 条件语句为 false
	4、case 条件语句为 true
	*/
}

/*
	[Type Switch]
	switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
*/ 
func GoConditionG(){
	var x interface{}
	switch i := x.(type) {
		case nil:  
			fmt.Printf(" x 的类型 :%T",i)                
		case int:  
			fmt.Printf("x 是 int 型")                      
		case float64:
			fmt.Printf("x 是 float64 型")          
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")                      
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )      
		default:
			fmt.Printf("未知型")    
   }  
}