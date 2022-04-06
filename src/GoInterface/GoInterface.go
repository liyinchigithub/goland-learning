package GoInterface

import (
	"log"
	"testing"
)

type Student interface {
	run()
	say()
}

type School struct {
	Student // 结构体继承接口
}

// 结构体实现接口方法
func (s *School) run() {
	log.Println("我会跑步")
}

// 结构体实现接口方法
func (s *School) say() {
	log.Println("我会说")
}

func Test(t *testing.T) {
	// 实例化结构体
	s := new(School)
	s.run()
	s.say()
}
