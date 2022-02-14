# goland-learning

>https://www.runoob.com/go/go-tutorial.html

## 安装

### 环境变量配置

安装Go语言需要配置的环境变量有GOROOT、GOPATH和Path

#### 【Windows】

假设Go安装路径为/usr/local/go

(1)配置GOROOT
GOROOT的变量值即为GO的安装目录/usr/local/go
(2)配置GOPATH
GOPATH的变量值即为存储Go语言项目的路径/usr/local/go/project
(3)配置Path
Path中有其他安装程序的配置信息，这里再增加一个GO的bin目录/usr/local/go/bin
(4)验证是否配置成功
```shell
go env
```
(5)通过命令查看安装的Go的版本
```shell
go version
```

#### 【linux】

1.下载二进制包

```shell
wget https://golang.google.cn/dl/go1.17.6.linux-amd64.tar.gz
```

2.将下载的二进制包解压至 /usr/local目录
```shell
cd /usr/local
tar -zxvf go1.17.6.linux-amd64.tar.gz
```

3.将 /usr/local/go/bin 目录添加至 PATH 环境变量

(1)编辑配置文件
```shell
cd ~
vim /etc/profile
```
(2)加入内容
```shell  
# goland
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/gopath/
export PATH=$PATH:$GOROOT/bin
```
>在gopath文件夹下新建src、bin、pkg三个文件夹
(3)刷新
```shell
source /etc/profile
```


#### 【mac】

[下载pkg安装包](https://golang.google.cn/dl/)

```shell
# goland
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/gopath/
export PATH=$PATH:$GOROOT/bin
```
>在gopath文件夹下新建src、bin、pkg三个文件夹

```shell
source .bash_profile
```

[hombrew安装方式]
```shell
brew install go
```


## 运行脚本

```shell
go run ./src/test.go
```
## 使用 go build 命令来生成二进制文件

```shell
go build ./src/hello.go
./hello
```

## 安装vscode go 插件

```shell
go get -v github.com/sqs/goreturns
```


# 依赖包管理


## 查看已安装依赖包列表
```shell
go list -m
```
## 国内镜像

### 【linux、mac】
```shell
# 官方
go env -w  GOPROXY=https://goproxy.io
# 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 七牛
go env -w  GOPROXY=https://goproxy.cn
# 确认镜像是否成功
go env | grep GOPROXY
```

### 【windows】
可以在 PowerShell 中设置：：

#### 启用 Go Modules 功能
```
$env:GO111MODULE="on"


## 安装

```shell
```

>https://go.dev/blog/using-go-modules


# 单元测试

使用testing包

```goland
package unit_test

// go单元测试
import (
	"testing"
)

/*
	[源文件]
	以_test 结尾：xxx_test.go
	[测试方法名]
	以Test开头：func TestName(t *testing.T){…}	方法名第一个字母不能是小写字母。
	[捕获异常]
	在这些函数中，使用 Error、Fail 或相关方法来发出失败信号。
	要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 TestXxx 函数，如上所述。
	将该文件放在与被测试文件相同的包中。该文件将被排除在正常的程序包之外，但在运行 go test 命令时将被包含。
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
	cd src/unit-test/
	go test -v -run=TestA unit_test.go //测试执行TestA 方法
	go test -v -run=TestB unit_test.go //测试执行TestB 方法
	go test -v unit_test.go //测试执行unit_test.go文件中的所有方法
	go test . // 测试文件夹下所有
*/

```




# 常见问题

* 1.执行go list 提示”go list -m: not using modules“
解决办法：
go依赖模块是通过[go module](https://go.dev/blog/using-go-modules)管理

* 2.执行go env 提示”go: GOPATH entry is relative; must be absolute path: "=/Users/liyinchi/gopath".For more details see: 'go help gopath'“
解决办法：环境配置文件中gopath和goroot不能一样，在gopath文件夹下新建三个文件夹src、pkg、bin

* 3.执行go get github.com/xuri/excelize/v2 提示”go: could not create module cache: mkdir /usr/local/go/gopath/pkg/mod: permission denied“
解决办法：
先执行mkdir /usr/local/go/gopath/pkg/mod
再执行sudo go getgithub.com/xuri/excelize/v2

* 4.

5.