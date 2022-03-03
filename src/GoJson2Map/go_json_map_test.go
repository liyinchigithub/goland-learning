package GoJson2Map

import(
	"fmt"
	"testing"
	"encoding/json"
)

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
		var event map[string]interface{}
		if err := json.Unmarshal(msg, &event); err != nil {
			panic(err)
		}
	
		fmt.Println(event)
}
