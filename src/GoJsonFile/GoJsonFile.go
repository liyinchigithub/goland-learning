package GoJsonFile

//  参考：https://www.jianshu.com/p/681b1c3b5a75

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

/*
配置	描述
alias	连接标识
driver	数据库驱动
database	数据库名称
user	数据库用户名
password	数据库密码
charset	数据库字符集
host	数据库主机地址
port	数据库端口号
max_idle_conns	最大空闲连接数
max_open_conns	最大打开连接数
*/

//load 加载JSON文件解析为字典映射
func load(filepath string) map[string]interface{} {
	//读取文件
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	//转化为字符串
	str := string(buf[:])
	//去除注释
	reg := regexp.MustCompile(`/\*.*\*/`)
	str = reg.ReplaceAllString(str, "")
	//反序列化
	obj := make(map[string]interface{})
	bs := []byte(str)
	err = json.Unmarshal(bs, &obj)
	if err != nil {
		panic(err)
	}
	return obj
}
