package GoStruct

import (
	"testing"
	"GoStruct" // 自定义封装模块放入/usr/local/go/src/下
)

func TestGoFunction(t *testing.T) {
	t.Log("开始测试")
	GoStruct.GoStructA()
	GoStruct.GoStructB()
	t.Log("测试完毕")
	// GoStructB()
	// GoStructC()
	// GoStructD()
	
}

/*
	运行脚本：
	cd /src/GoStruct
	go test -v go_struct_test.go
*/
