# goland-learning

[![goland](https://img.shields.io/badge/goland-1.17.7-green.svg)](https://github.com/golang/go) 


## 安装

### 环境变量配置

安装Go语言需要配置的环境变量有GOROOT、GOPATH和Path

#### 【Windows】

假设Go安装路径为/usr/local/go

* (1)配置GOROOT
GOROOT的变量值即为GO的安装目录/usr/local/go
* (2)配置GOPATH
GOPATH的变量值即为存储Go语言项目的路径/usr/local/go/project
* (3)配置Path
Path中有其他安装程序的配置信息，这里再增加一个GO的bin目录/usr/local/go/bin
* (4)验证是否配置成功
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

查看GOROOT环境变量
```shell
echo $GOROOT
```

[hombrew安装方式]
```shell
brew install go
```
#### 模块管理

>https://go.dev/doc/tutorial/create-module

```shell
go mod init
```
执行命令go mod init在当前目录下生成一个go.mod文件，执行这条命令时，当前目录不能存在go.mod文件。如果之前生成过，要先删除；
go.mod 文件来跟踪代码的依赖关系。
该文件仅包含模块的名称和代码支持的 Go 版本。
但是当您添加依赖项时，go.mod 文件将列出您的代码所依赖的版本。
这使构建保持可重复性，并使您可以直接控制要使用的模块版本

mod.go文件内容如下：

```go
module demo

go 1.17

require github.com/360EntSecGroup-Skylar/excelize v1.4.1

require (
	github.com/bunsenapp/go-selenium v0.1.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	github.com/xuri/excelize/v2 v2.5.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/text v0.3.6 // indirect
)

```

[注意]Linux、mac，安装完依赖后引入第三方模块还是会提示，可能会报/usr/local/go/gopath/pkg/mod/ 没有权限

解决办法：
```shell
sudo chmod 777 /usr/local/go/gopath/pkg/mod/
```

## 安装依赖

如果有环境已有，直接执行安装依赖

```shell
go mod tidy
```
执行go mod tidy命令，它会添加缺失的模块以及移除不需要的模块。
执行后会生成go.sum文件(模块下载条目)。添加参数-v，例如:go mod tidy -v可以将执行的信息，即删除和添加的包打印到命令行

## 其他命令

### go mod verify
执行命令go mod verify来检查当前模块的依赖是否全部下载下来，是否下载下来被修改过。
如果所有的模块都没有被修改过，那么执行这条命令之后，会打印all modules verified。

### go mod vendor
执行命令go mod vendor生成vendor文件夹，该文件夹下将会放置你go.mod文件描述的依赖包，文件夹下同时还有一个文件modules.txt，它是你整个工程的所有模块。
在执行这条命令之前，如果你工程之前有vendor目录，应该先进行删除。
同理go mod vendor -v会将添加到vendor中的模块打印出来；


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
```

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
	cd src/GoUnitTest/
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

解决办法：

环境配置文件中gopath和goroot不能一样，在gopath文件夹下新建三个文件夹src、pkg、bin

* 3.执行go get github.com/xuri/excelize/v2 提示”go: could not create module cache: mkdir /usr/local/go/gopath/pkg/mod: permission denied“

解决办法：

先执行mkdir /usr/local/go/gopath/pkg/mod
再执行sudo go getgithub.com/xuri/excelize/v2

* 4.安装依赖报错在镜像上面没有找到
执行 go get xxxx
go get: github.com/electricbubble/gidevice@v0.6.0: reading https://mirrors.aliyun.com/goproxy/github.com/electricbubble/gidevice/@v/v0.6.0.info: 404 Not Found

解决办法：切换镜像到官方镜像或七牛

```shell
# 官方
go env -w  GOPROXY=https://goproxy.io
w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 七牛
go env -w  GOPROXY=https://goproxy.cn
# 确认镜像是否成功
go env | grep GOPROXY
```

* 5.格式化代码

```shell
go fmt .
```

* 6.模块导入

--包A

--文件a

----包B

----文件b

在文件b中引入包a的函数

<img width="826" alt="image" src="https://user-images.githubusercontent.com/19643260/161190715-4523ce85-10d9-4995-97af-a887ae361899.png">


```shell

```



* 7.单元测试testing 控制台调试输出

testing使用t.log() 也可以使用fmt.println()，但不能直接go test .  而是需要指定执行文件名

```shell
package test

import (
	"app/api"
	"fmt"
	"testing"
)

func TestGoHttpClientA(t *testing.T) {
	api.GentlemanSampleRequest()
	t.Log("测试中...string")
	fmt.Println("测试中...string")
}

/*
	命令行执行，运行指定文件名、文件中某个函数：
	go test -v go_test.go
	go test -v -run=TestGoHttpClientA go_test.go
*/

```

或

```shell
sudo /usr/local/go/bin/go test -timeout 30s
```


```shell
cd /goland-learning/src/GoExcel/test
go test -v -run=TestGoExcelB .
```

* 8.

```shell

```


##  参考教程

>https://www.runoob.com/go/go-tutorial.html

>https://gfw.go101.org/article/101.html