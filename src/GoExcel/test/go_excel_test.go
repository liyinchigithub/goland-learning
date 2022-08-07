package GoExcel

import (
	"fmt"
	"src/GoExcel"
	"testing"
)

func TestGoExcelA(t *testing.T) {
	//
	// GoExcel.TestCreateExcel()
}

func TestGoExcelB(t *testing.T) {
	GoExcel.TestReadExcel()
}

func TestGoExcelC(t *testing.T) {
	// 声明集合
	map_data := map[string][]string{}
	// 声明数组并初始化
	names := []string{"张三", "李四", "王五"}
	fmt.Printf("names:%v\n", names)
	// 声明数组并初始化
	address := []string{"泉州市丰泽区", "厦门市集美区", "福州市金山区"}
	fmt.Printf("address:%v\n", address)
	// 声明数组并初始化
	numbers := []string{"2020001", "2020002", "2020003"}
	fmt.Printf("numbers:%v\n", numbers)

	map_data["names"] = names
	map_data["address"] = address
	map_data["numbers"] = numbers
	t.Logf("map_data:%v\n", map_data) // map_data:map[address:[泉州市丰泽区 厦门市集美区 福州市金山区] names:[张三 李四 王五] numbers:[123123 234234 4534534]]

	// 追加写入excel
	GoExcel.TestAppendExcel("Book1.xlsx", "Sheet1", map_data)
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
