package GoOS

import(
	"os"
	"fmt"
)

func GoOS() {
	var arr = os.Args
	fmt.Println(arr)
	if string(arr[1]) == "ok" {
		fmt.Println(`是OK%d`, arr[1])
		os.Exit(0) //main方法退出
	} else {
		fmt.Println(arr)
		fmt.Println(`不是OK%d`, arr[1])
	}
}

/* 

	运行单元测试
	sudo /usr/local/go/bin/go test -timeout 30s -run ^TestGoOS$
*/