package test

import (
	"fmt"
)

/*
	[延迟执行]
	defer执行逻辑：
	当程序执行到一个defer时，不会立即执行defer后面的语句，而是将defer后的语句压入到一个专门存储defer语句的栈中，当程序执行到return、panic、recover、结束程序时，会依次从栈中取出defer语句并执行。
*/

func GoTestDefer() {
	defer A() //  1
	defer B() // 2
	defer C() // 3
	/*
		输出：
			C: 3
			B: 2
			A: 1
	*/

}

func A() {
	fmt.Printf("A: %d\n", 1)
}
func B() {
	fmt.Printf("B: %d\n", 2)
}
func C() {
	fmt.Printf("C: %d\n", 3)
}
