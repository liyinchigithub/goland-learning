package test

import (
	GoMath "src/GoMath"
	"testing"
)

/*
	golang生成随机数可以使用math/rand包
*/

func TestMathRandom(t *testing.T) {
	t.Log("start test")
	// GoMath.TestMathRandomA1()
	// GoMath.TestMathRandomA2()// 没有seed，结果都是一样的
	GoMath.TestMathRandomB()
	// GoMath.TestMathRandomC()
	// GoMath.TestMathRandomD1()
	// GoMath.TestMathRandomD2()
	// GoMath.TestMathRandomE()
	// GoMath.TestMathRandomF()
	t.Log("end test")
}

/*
	命令行执行：
	cd /goland-learning/src/GoMath/test
	go test -v -run TestMathRandom
	go test -v -run TestMathRandom go_math_rand_test.go
	go test -v go_math_rand_test.go
*/
