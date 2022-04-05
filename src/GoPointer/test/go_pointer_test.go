package test

import (
	"testing"
	GoPointer "src/GoPointer"
)

func TestGoPointer(t *testing.T){
	t.Log("指针变量和指针地址:")
	GoPointer.GoPointerB()
	GoPointer.GoPointerExchange()
}

/*
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -timeout 30s -run TestGoExcelA
	[单元测试文件名以_test结尾]
	[函数大写Test开头]
*/