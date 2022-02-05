# goland-learning

>https://www.runoob.com/go/go-tutorial.html

## 环境安装

>https://golang.google.cn/dl/


1.Go语言环境变量配置
安装Go语言需要配置的环境变量有GOROOT、GOPATH和Path

例如：Go安装路径为/soft/Go/

(1)配置GOROOT
GOROOT的变量值即为GO的安装目录/soft/Go/
(2)配置GOPATH
GOPATH的变量值即为存储Go语言项目的路径/soft/Go/project
(3)配置Path
Path中有其他安装程序的配置信息，这里再增加一个GO的bin目录/soft/Go/bin
(4)验证是否配置成功
```shell
go env
```
(5)通过命令查看安装的Go的版本
```shell
go version
```

```
export GOROOT=/usr/local/Cellar/go/1.14.2_1/libexec
export GOARCH=amd64
export GOOS=darwin
export GOPATH=/Users/liyinchi/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
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