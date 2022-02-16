package Go99

import (
	"testing"
)

func TestGo99(t *testing.T){
	// t.Log("99乘法表")
	Test99()
}

/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -timeout 30s -run TestGo99
	[单元测试文件名以_test结尾]
	[函数大写Test开头]

	单元测试入参使用数据驱动，不是用命令行，os.args仅引用于go run 脚本.go 入参...
*/