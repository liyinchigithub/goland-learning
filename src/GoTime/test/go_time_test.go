package GoTime

import (
	"testing"
	GoTime "src/GoTime"
)

func TestGoFunction(t *testing.T) {
	GoTime.GoTimeA()
	GoTime.GoTimeB()
	GoTime.GoTimeC()
	GoTime.GoTimeD()
	/*
	输出：
		2020-04-26 17:51:46 +0800 CST
		2020-04-26 17:51:46
		2022/04/03 14:02:39
		2022-04-03 14:02:39.573028 +0800 CST m=+0.000582763
		2022-04-03 14:02:39当前时间戳： 1648965759
		当前纳秒时间戳： 1648965759573040000
	*/
}

/*
	运行脚本
	cd /src/GoTime/test
	go test -v go_time_test.go
*/
