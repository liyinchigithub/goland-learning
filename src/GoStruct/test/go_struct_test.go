package GoStruct

import (
	"src/GoStruct"
	"testing"
)

func TestGoFunction(t *testing.T) {
	t.Log("开始测试")
	// GoStruct.GoStructA()
	// GoStruct.GoStructB()
	// GoStruct.GoStructB()
	// GoStruct.GoStructC()
	// GoStruct.GoStructD()
	// GoStruct.GoStructE()
	GoStruct.StructInitType01()
	t.Log("测试完毕")

}

/*
	运行脚本：
	cd /src/GoStruct
	go test -v go_struct_test.go
*/
