package test

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	str := pwd()
	log.Printf("%v\n", str)

	str = path.Join(str, "../database.json")
	log.Printf("%v\n", str)
}

//pwd 获取当前文件所在路径
func pwd() string {
	str, err := os.Getwd() // 当前的路径
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", str) // /Users/mac/go/src/GoJsonFile/test
	str, err = filepath.Abs(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("str为%v\n", str)
	str = strings.Replace(str, "\\", "/", -1)
	return str
}
