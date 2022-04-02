package test

import  (
	"testing"
	"src/GoCondition"
)

func TestGoCondition(t *testing.T){
	GoCondition.GoConditionA()
	GoCondition.GoConditionB()
	GoCondition.GoConditionC()
	GoCondition.GoConditionD()
	GoCondition.GoConditionE()
	GoCondition.GoConditionF()
	GoCondition.GoConditionG()
}

/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -timeout 30s -run ^函数名$
	sudo /usr/local/go/bin/go test -timeout 30s 文件名.go
	（3）go test -v go_condition_test.go
	[单元测试文件名以_test结尾]
	[函数大写Test开头]
*/