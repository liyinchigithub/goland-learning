package GoOS

import "testing"


func TestGoOS(t *testing.T){
	GoOS()
	t.Log("done")
}

/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	/usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本
	/usr/local/go/bin/go test -timeout 30s -run TestGoOS
	
	[单元测试文件名以_test结尾]
	[函数大写Test开头]
*/