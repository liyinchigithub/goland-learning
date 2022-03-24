package GoHttpClient

import (
	"fmt"
  // gentleman内置了将近 20 个插件，有身份认证相关的auth、有cookies、有压缩相关的compression、有代理相关的proxy、有重定向相关的redirect、有超时相关的timeout、有重试的retry、有服务发现的consul等等
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/mux"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
  "gopkg.in/h2non/gentleman.v2/plugins/headers" // header插件用于在发送请求前添加一些通用的首部，如 APIKey；或者删除一些自动加上的首部，如User-Agent。一般header插件应用在cli对象上
  // "gopkg.in/h2non/gentleman.v2/plugins/query" // HTTP 请求通常会在 URL 的?后带上查询字符串（query string），gentleman的内置插件query可以很好的管理这个信息
)

/*
 	[API]
	https://pkg.go.dev/gopkg.in/h2non/gentleman.v2
  [参考]
  https://zhuanlan.zhihu.com/p/127398199
*/

func GentlemanSampleRequest() {
	// Create a new client
	cli := gentleman.New() // 创建一个 HTTP 客户端cli

	// Define the Base URL
	cli.URL("http://httpbin.org/post") // 设置要请求的 URL 基础地址
  cli.Use(headers.Set("Content-type", "application/json"))
	// Create a new request based on the current client
	req := cli.Request() // 创建一个请求对象req

	// Method to be used
	req.Method("POST") // 设置请求方法

	// req.Path("/api/breeds/image/random") // 设置请求的路径，基于前面设置的 URL

	// req.SetHeader("Client", "gentleman") // 设置请求首部（Header），上面代码设置首部Client为gentleman

	// Define the JSON payload via body plugin
	data := map[string]string{"foo": "bar"}
	req.Use(body.JSON(data))

	// Perform the request
	res, err := req.Send() // 发送请求，获取响应对象res
	//
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	//
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	// 响应状态码
	fmt.Printf("Status: %d\n", res.StatusCode)
	// 响应内容
	fmt.Printf("Body: %s", res.String())
}

func GentlemanSendJsonBody() {
	// Create a new client
	cli := gentleman.New()

	// Define the Base URL
	cli.URL("http://httpbin.org/post")

	// Create a new request based on the current client
	req := cli.Request()

	// Method to be used
	req.Method("POST")

	// Define the JSON payload via body plugin
	data := map[string]string{"foo": "bar"}
	req.Use(body.JSON(data))

	// Perform the request
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())

}

func GentlemanSenXMLBody() {
}
func GentlemanCompositionViaMultiplexer() {
	// Create a new client
	cli := gentleman.New()

	// Define the server url (must be first)
	cli.Use(url.URL("http://httpbin.org"))

	// Create a new multiplexer based on multiple matchers
	mx := mux.If(mux.Method("GET"), mux.Host("httpbin.org"))

	// Attach a custom plugin on the multiplexer that will be executed if the matchers passes
	mx.Use(url.Path("/headers"))

	// Attach the multiplexer on the main client
	cli.Use(mx)

	// Perform the request
	res, err := cli.Request().Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
