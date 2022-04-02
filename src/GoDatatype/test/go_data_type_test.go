package test

import (
	"testing"
	"src/GoDatatype"
)


func TestGoDatatype(t *testing.T){
	// 同一个包下，可以互相直接饮用包中的函数，无需引入
	GoDatatype.GoDatatype()
}

/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -timeout 30s -run ^GoDatatype$
	sudo /usr/local/go/bin/go test -timeout 30s go_data_type_test.go
	（3）go test -v go_condition_test.go
	[单元测试文件名以_test结尾]
	[函数大写Test开头]
*/