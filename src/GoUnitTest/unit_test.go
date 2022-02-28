package GoUnitTest

// go单元测试
import (
	"testing"
)

/*
	[源文件]
			文件名必须以_test 结尾：xxx_test.go
	[方法名]
			以Test开头：func TestName(t *testing.T){…} 方法名第一个字母不能是小写字母。
	[捕获异常]
			在这些函数中，使用 Error、Fail 或相关方法来发出失败信号。
			要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 TestXxx 函数，如上所述。
			将该文件放在与被测试文件相同的包中。
			该文件将被排除在正常的程序包之外，但在运行 go test 命令时将被包含。
			有关详细信息，请运行 go help test 和 go help testflag 了解
	[跳过测试]
			可以调用 *T 和 *B 的 Skip 方法，跳过该测试或基准测试
			t.Skip("skipping test in short mode.")
*/

func TestA(t *testing.T) {
	Fib(4)
	// 日志
	t.Log("单元测试TestA")
}

func TestB(t *testing.T) {
	// 日志
	t.Log("单元测试TestB")
	// fmt.println("单元测试TestB") // 单元你测试中不使用这个会异常报错 ./unit_test.go:34:2: cannot refer to unexported name fmt.println
}

// 递归函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-1)
}

func TestFib(t *testing.T) {
	var (
		// 定义变量
		in = 7
		// 定义变量
		expected = 13
	)
	// 定义变量
	actual := Fib(in) // 变量actual接收函数Fib的返回值
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}

/*
	[命令行运行]
	cd src/GoUnitTest/
	go test -v -run=TestA unit_test.go //测试执行TestA 方法
	go test -v -run=TestB unit_test.go //测试执行TestB 方法
	sudo /usr/local/go/bin/go test -v -timeout 30s -run TestB
	sudo /usr/local/go/bin/go test unit_test.go -timeout 30s 
	go test -v unit_test.go //测试执行unit_test.go文件中的所有方法
	go test . // 测试文件夹下所有
	go test  -v -timeout 30s
*/


// 数据驱动
func TestTableFib(t *testing.T) {
    var fibTests = []struct {
        in       int // input
        expected int // expected result
    }{
        {1, 1},
        {2, 1},
        {3, 2},
        {4, 3},
        {5, 5},
        {6, 8},
        {7, 13},
    }

    for _, v := range fibTests {
        actual := Fib(v.in)
        if actual != v.expected {
            t.Errorf("Fib(%d) = %d; expected %d", v.in, actual, v.expected)
        }
    }
}