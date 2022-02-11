package test_go_excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

/*
	[go创建读取excel文件]
	Excelize 是 Go 语言编写的一个用来操作 Office Excel 文档类库，基于 ECMA-376 Office OpenXML 标准。
	可以使用它来读取、写入 XLSX 文件，相比较其他的开源类库，Excelize 支持写入原本带有图片(表)的文档，还支持向 Excel 中插入图片，并且在保存后不会丢失图表样式。
	[注意事项]
	使用最新版本 Excelize 要求您使用的 Go 语言为 1.15 或更高版本
	[excelize官方项目]
	https://github.com/360EntSecGroup-Skylar/excelize
	[excelize官方API]
	https://xuri.me/excelize/zh-hans/
	[博文]
	https://blog.csdn.net/weixin_43456598/article/details/100678077
	https://studygolang.com/articles/20337
*/

// 创建excel
func createExcel() {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// 读取excel
func readExcel() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
