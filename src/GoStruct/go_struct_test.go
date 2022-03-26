package GoStruct

import (
	"testing"
	"GoStruct" // 自定义封装模块放入/usr/local/go/src/下
)

func TestGoFunction(t *testing.T) {
	GoStruct.GoStructA()
	GoStruct.GoStructB()
	// GoStructB()
	// GoStructC()
	// GoStructD()
	
}

/*
	运行脚本：
	cd /src/GoStruct
	go test -v go_struct_test.go
*/
