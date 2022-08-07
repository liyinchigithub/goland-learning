package GoExcel

import (
	"fmt"
	"reflect"

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
func TestCreateExcel() {
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
func TestReadExcel() {
	// 打开文件
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 捕获异常
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "A2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("cell", cell)
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

// func main() {
// readEecel()
// writeEecel()
// appendEecel()
// }

//  覆盖写入
func TestWriterEecel(sheetName string, cell string, writerContent string) {
	//  实例化结构体对象f，并调用结构体方法NewFile新建文件
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	// 设置单元格的值
	f.SetCellValue("Sheet1", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

/*
* @methods TestAppendExcel
* @description 追加数据
* @params fileName string	文件名
* @params activeSheet string	Sheet名
* @params writerData map[string][]string{}	写入数据
 */
func TestAppendExcel(fileName string, activeSheet string, writerData map[string][]string) {
	//  文件名
	// fileName := "Book1.xlsx"
	//  工作表
	// activeSheet := "Sheet1"
	//  打开文件
	f, _ := excelize.OpenFile(fileName)
	// Get all the rows in the Sheet1
	//  获取所有行的数据，按行获取
	rows, err := f.GetRows(activeSheet)
	// 行数据
	fmt.Printf("rows_content：%v\n", rows)
	//  行数
	fmt.Println("rows_len:", len(rows))
	//  捕获异常
	if err != nil {
		fmt.Println(err)
		return
	}
	//  获取所有列的内容 按列获取
	cols, err := f.GetCols(activeSheet)
	//  捕获异常
	if err != nil {
		fmt.Println(err)
		return
	}
	// 列数据
	fmt.Printf("cols_content：%v\n", cols)
	//  列数
	fmt.Println("cols_len:", len(cols))
	//  行高
	lineHeight, err := f.GetRowHeight(activeSheet, 1)
	fmt.Println("lineHeight:", lineHeight)
	//  捕获异常
	if err != nil {
		fmt.Println(err)
		return
	}
	//  行级别？？？
	lineLevel, err := f.GetRowOutlineLevel(activeSheet, 2)
	fmt.Println("lineLevel:", lineLevel)
	//  捕获异常
	if err != nil {
		fmt.Println(err)
		return
	}
	// 遍历写入数据(遍历map元素是字符串数组)
	for _, arr_v := range writerData["names"] {
		fmt.Println("writerData['names']", arr_v)
		//  判断数据类型
		fmt.Println("writerData['names'] arr_v type:", reflect.TypeOf(arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("A%d", len(rows)+1), fmt.Sprintf("%v", arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("A%d", len(rows)+2), fmt.Sprintf("%v", arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("A%d", len(rows)+3), fmt.Sprintf("%v", arr_v))
	}
	for _, arr_v := range writerData["address"] {
		fmt.Println("writerData['address']", arr_v)
		//  判断数据类型
		fmt.Println("writerData['address'] arr_v type:", reflect.TypeOf(arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("B%d", len(rows)+1), fmt.Sprintf("%v", arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("B%d", len(rows)+2), fmt.Sprintf("%v", arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("B%d", len(rows)+3), fmt.Sprintf("%v", arr_v))
	}
	for i, arr_v := range writerData["numbers"] {
		fmt.Println("writerData['numbers']", arr_v)
		//  判断数据类型
		fmt.Println("writerData['numbers'] arr_v type:", reflect.TypeOf(arr_v))
		f.SetCellValue(activeSheet, fmt.Sprintf("C%d", len(rows)+1), fmt.Sprintf("%v", arr_v[i]))
		f.SetCellValue(activeSheet, fmt.Sprintf("C%d", len(rows)+2), fmt.Sprintf("%v", arr_v[i]))
		f.SetCellValue(activeSheet, fmt.Sprintf("C%d", len(rows)+3), fmt.Sprintf("%v", arr_v[i]))
		fmt.Println("arr_v[i]:", arr_v[i])
	}
	//  单元格值设置
	// f.SetCellValue(activeSheet, fmt.Sprintf("A%d", len(rows)+1), fmt.Sprintf("aa%d", len(rows)+1))
	// f.SetCellValue(activeSheet, fmt.Sprintf("B%d", len(rows)+1), fmt.Sprintf("bb%d", len(rows)+1))
	// f.SetCellValue(activeSheet, fmt.Sprintf("C%d", len(rows)+1), fmt.Sprintf("cc%d", len(rows)+1))
	//  保存文件，并捕获异常
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}

func TestReadEecel() {
	f, err := excelize.OpenFile("Book1.xlsx")
	// f, err := excelize.OpenFile("信息系统暴露面模板_20211027.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//  捕获异常
	defer func() {
		// 关闭文件读写
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//fmt.Println(cell)

	// 获取 Sheet1 上所有单元格（所有行数据）
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//  遍历
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	fmt.Println("===========================================")

	res := ArrayTwoStringGroupsOf(rows, 3)
	fmt.Println(res)

}

//
func ArrayTwoStringGroupsOf(arr [][]string, num int64) [][][]string {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

//
func ArrayInGroupsOf(arr []int, num int64) [][]int {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}
