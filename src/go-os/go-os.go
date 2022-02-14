package test_go_os

import "os"

func test_go_os() {
	var arr = os.Args
	fmt.Println(arr)
	if arr[0] == "ok" {
		os.Exit(0) //main方法退出
	} else {
		println(`不是OK%v`, arr[1])
	}
}
