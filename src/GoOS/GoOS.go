package GoOS

import(
	"os"
	"fmt"
)

func GoOS() {
	var arr = os.Args
	fmt.Println(arr)
	if arr[0] == "ok" {
		os.Exit(0) //main方法退出
	} else {
		println(`不是OK%v`, arr[1])
	}
}

/* 

	运行单元测试
	/usr/local/go/bin/go test -timeout 30s -run ^TestGoOS$
*/