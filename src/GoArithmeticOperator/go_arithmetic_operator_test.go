package GoArithmeticOperator

import (
	"testing"
	"fmt"
)

func TestA(t *testing.T)  {
	var a = 10
	var b = 3
	t.Log(float64(a / b))
	fmt.Println(a / b) //3 都是int 返回int
	var c = 10.0
	t.Log(c / float64(b))
	// fmt.Println(c / float64(b)) //3.3333333333333335 //都是float 返回float
}


/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -timeout 30s -run ^GoArithmeticOperator$
	sudo /usr/local/go/bin/go test -timeout 30s go_arithmetic_operator_test.go
	[单元测试文件名以_test结尾]
	[函数大写Test开头]
*/