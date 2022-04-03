package GoJson2Map

import(
	"fmt"
	"testing"
	"encoding/json"
)

// Json转Map
func Test(t *testing.T)  {
	msg := []byte(`{
		"type": "UPDATE",
		"database": "blog",
		"table": "blog",
		"data": [
			{
				"blogId": "100001",
				"title": "title",
				"content": "this is a blog",
				"uid": "1000012",
				"state": "1"
			}
		]}`)
		// 定义一个map key为字符串 value为interface{} value可以接受任何类型
		var event map[string]interface{}
		// 解析json字符串
		if err := json.Unmarshal(msg, &event); err != nil {// &event是一个指针，指向一个map 引用变量event将受到变化影响
			panic(err)
		}
	
		fmt.Println(event)
}
