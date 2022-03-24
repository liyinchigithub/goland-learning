package GoTime

import (
	"fmt"
	"time"
)

/*
	[时间戳转换成日期字符串]
*/
func GoTimeA() {
	// unixTime: 1587888473
	unixTime := 1587894706
	timeObj := time.Unix(int64(unixTime), 0)
	fmt.Println(timeObj)
	var str = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str) //2020-04-26 17:51:46
}

/*
	[打印当前日期(直接获取)]
	时间类型有一个自带的方法Format进行格式化，
	需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
	而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）


	2006  年
	01  月
	02  日
	03  时   12小时制     写成：15   表示   24小时制
	04  分
	05  秒

*/
func GoTimeB() {
	// timeObj := time.Now()
	// var str = timeObj.Format("2006-01-02 03:04:05")
	// fmt.Println(str) //2020-04-26 05:44:24

	//12小时制
	// timeObj := time.Now()
	// var str = timeObj.Format("2006/01/02 03:04:05")
	// fmt.Println(str) //2020/04/26 05:44:47

	// 24小时制
	timeObj := time.Now()
	var str = timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(str) // 2020/04/26 17:48:53
}

/*
	[打印当前日期(单独取年，月，日，时，分，秒)]
*/

func GoTimeC() {
	timeObj := time.Now()
	fmt.Println(timeObj) //2020-04-26 17:32:25.9639049 +0800 CST m=+0.004000301

	// 2020-04-26 17:32:25
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	// fmt.Printf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second) //2020-4-26 17:35:07

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second) //2020-04-26 17:36:04

	//注意：%02d 中的 2 表示宽度，如果整数不够 2 列就补上 0
}

/*
	[获取时间戳]
	时间戳是自 1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数。它也被称为 Unix 时 间戳（UnixTimestamp
*/
func GoTimeD() {
	timeObj := time.Now()

	unixtime := timeObj.Unix()      //获取当前的时间戳 （秒）
	fmt.Println("当前时间戳：", unixtime) //当前时间戳： 1587894706

	unixNatime := timeObj.UnixNano()    //纳秒时间戳
	fmt.Println("当前纳秒时间戳：", unixNatime) //当前时间戳：  1587894791217129300
}

/*
	[日期间隔]
*/
func GoTimeG() {
	/*
	   1、time包中定义的时间间隔类型的常量如下：
	       const (
	           Nanosecond  Duration = 1
	           Microsecond          = 1000 * Nanosecond
	           Millisecond          = 1000 * Microsecond
	           Second               = 1000 * Millisecond
	           Minute               = 60 * Second
	           Hour                 = 60 * Minute
	       )
	*/
	// fmt.Println(time.Millisecond) //1毫秒
	// fmt.Println(time.Second)      //1秒

	/*
	   2、时间操作函数
	*/

	var timeObj = time.Now()
	fmt.Println(timeObj)
	timeObj = timeObj.Add(time.Hour)
	fmt.Println(timeObj)

	/*
	   2020-04-26 18:15:05.4612997 +0800 CST m=+0.005000201
	   2020-04-26 19:15:05.4612997 +0800 CST m=+3600.005000201
	*/

}

/*
	[定时器]
*/
func GoTimeH() {
	// // time.Now()
	// ticker := time.NewTicker(time.Second)
	// // ticker.C
	// for t := range ticker.C {
	//     fmt.Println(t)
	// }

	// time.Now()
	// ticker := time.NewTicker(time.Second)
	// n := 5
	// for t := range ticker.C {
	//     n--
	//     fmt.Println(t)
	//     if n == 0 {
	//         ticker.Stop() //终止这个定时器继续执行
	//         break
	//     }

	// }

	//休眠方法
	// fmt.Println("aaa")
	// time.Sleep(time.Second)
	// fmt.Println("aaa2")
	// time.Sleep(time.Second)
	// fmt.Println("aaa3")
	// time.Sleep(time.Second * 5)
	// fmt.Println("aaa4")

	for {
		time.Sleep(time.Second)
		fmt.Println("我在定时执行任务")
	}
}

/*
	[时间戳和日期相互转换]
*/
func GoTimeI() {
	//输出完整时间
	fmt.Println(time.Now()) //2021-04-28 16:08:56.138526 +0800 CST m=+0.020943201
	//输出秒级时间戳
	fmt.Println(time.Now().Unix()) // 1619597336
	//时间戳转换成时间字符串
	shijian := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println(shijian) //2021-04-28 16:10:48
	//时间字符串转换成时间戳
	str := "2021-03-19 15:04:05"
	std, _ := time.Parse("2006-01-02 15:04:05", str)
	fmt.Println(std)        //2021-03-19 15:04:05 +0000 UTC
	fmt.Println(std.Unix()) //1616166245
}
