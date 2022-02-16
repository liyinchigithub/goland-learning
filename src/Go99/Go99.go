package Go99

import "fmt"

func Test99()  {
	count := 9
	for i := 1; i < count; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d*%d=%d ",i, j, i*j) // 格式化数据要使用[printf]
		}
		 // 手动生成回车
		 fmt.Println()
	}
	
}

