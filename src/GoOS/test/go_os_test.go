package GoOS

import (
	"testing"
	GoOS "src/GoOS"
)


func TestGoOS(t *testing.T){
	GoOS.GoOS()
	t.Log("done")
}

/* 
	运行脚本：
	[注意]
	命令行参数只能使用
	main.go  go run main.go "ok"
	不能用单元测试go test -run go_os_test.go
*/