package GoExcel

import (
	"testing"
	"src/GoExcel"
)
func TestGoExcelA(t *testing.T){
	// 
	GoExcel.TestCreateExcel();
}

func TestGoExcelB(t *testing.T){
	GoExcel.TestReadExcel();
}

/* 
	[运行单元测试]
	（1）运行文件夹下所有单元测试脚本
	sudo /usr/local/go/bin/go test -timeout 30s
	（2）运行指定单元测试脚本中的函数
	sudo /usr/local/go/bin/go test -timeout 30s -run TestGoExcelA
	（3）go test -v go_excel_test.go
	[单元测试文件名以_test结尾]
	[函数大写Test开头]

	单元测试入参使用数据驱动，不是用命令行，os.args仅引用于go run 脚本.go 入参...
*/