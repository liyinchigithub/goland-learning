package GoStrings

import "testing"


func TestGoStrings(t *testing.T){
	// GoStringsContains() 
	// GoStringsSplit()
	// GoStringsJoin()
	// GoStringsJoinArray()
	StringRune()
	t.Log("done")
}

/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -v -timeout 30s -run GoStrings
	[单元测试文件名以_test结尾]
	[函数大写Test开头]
*/