package GoOS

import(
	"os"
	"fmt"
)

// 接收命令行参数
func GoOS() {
	var arr = os.Args
	fmt.Println(arr)
	if string(arr[1]) == "ok" {
		fmt.Printf(`是OK%d`, arr[1])
		os.Exit(0) //main方法退出
	} else {
		fmt.Println(arr)
		fmt.Printf(`不是OK%d`, arr[1])
	}
}

