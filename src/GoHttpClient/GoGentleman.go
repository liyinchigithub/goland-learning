package GoHttpClient

import (
	"fmt"
	"os"
	// gentleman内置了将近 20 个插件，有身份认证相关的auth、有cookies、有压缩相关的compression、有代理相关的proxy、有重定向相关的redirect、有超时相关的timeout、有重试的retry、有服务发现的consul等等
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/mux"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/headers" // header插件用于在发送请求前添加一些通用的首部，如 APIKey；或者删除一些自动加上的首部，如User-Agent。一般header插件应用在cli对象上
	"gopkg.in/h2non/gentleman.v2/plugins/url"
	// "gopkg.in/h2non/gentleman.v2/plugins/query" // HTTP 请求通常会在 URL 的?后带上查询字符串（query string），gentleman的内置插件query可以很好的管理这个信息
)

/*
 	[API]
	https://pkg.go.dev/gopkg.in/h2non/gentleman.v2
  [参考]
  https://zhuanlan.zhihu.com/p/127398199
*/

type Client struct {
	url string
}

func GentlemanSampleRequest(url string) {
	// Create a new client
	cli := gentleman.New() // 创建一个 HTTP 客户端cli

	// Define the Base URL
	cli.URL(url) // 设置要请求的 URL 基础地址
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

	/* 调试输出内容
	Status: 200
	Body: {
	"args": {},
	"data": "{\"foo\":\"bar\"}\n",
	"files": {},
	"form": {},
	"headers": {
		"Accept-Encoding": "gzip",
		"Content-Length": "14",
		"Content-Type": "application/json",
		"Host": "httpbin.org",
		"User-Agent": "gentleman/2.0.5",
		"X-Amzn-Trace-Id": "Root=1-6241c6d3-5583ba4a24bad7d60ae19b9e"
	},
	"json": {
		"foo": "bar"
	},
	"origin": "27.154.169.81",
	"url": "http://httpbin.org/post"
	}
	*/
}

/*
	content-type：application/json
*/
func GentlemanSendJsonBody(url string, obj map[string]string) {
	// 创建客户端（实例化）
	cli := gentleman.New()

	// 定义一个请求地址
	cli.URL(url)

	// 基于当前客户端创建一个新请求
	req := cli.Request()

	// 请求方法
	req.Method("POST")

	// 通过body插件，定义JSON有效负载
	data := obj
	req.Use(body.JSON(data))

	// 执行请求
	res, err := req.Send()
	// 判断是否请求异常
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	// 判断是否请求失败（请求成功 res.Ok=true）
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	// 响应状态码
	fmt.Printf("Status: %d\n", res.StatusCode)
	// 响应body
	fmt.Printf("Body: %s", res.String())

	/*
		调试输出内容：
		Status: 200
		Body: {
		"args": {},
		"data": "{\"username\":\"liyinchi\"}\n",
		"files": {},
		"form": {},
		"headers": {
			"Accept-Encoding": "gzip",
			"Content-Length": "24",
			"Content-Type": "application/json",
			"Host": "httpbin.org",
			"User-Agent": "gentleman/2.0.5",
			"X-Amzn-Trace-Id": "Root=1-6241c6d4-4851139f086fcb890b9f5746"
		},
		"json": {
			"username": "liyinchi"
		},
		"origin": "27.154.169.81",
		"url": "http://httpbin.org/post"
		}
	*/
}

/*
	content-type：application/xml
*/
func GentlemanSenXMLBody() {
	type User struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	cli := gentleman.New()
	cli.URL("http://httpbin.org/post")

	req := cli.Request()
	req.Method("POST")
	// 实例化结构体
	u := User{Name: "dj", Age: 18}
	// 使用 body.XML插件
	req.Use(body.XML(u))

	// 执行请求
	res, err := req.Send()
	// 判断是否请求异常
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	// 判断是否请求失败（请求成功 res.Ok=true）
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	// 响应状态码
	fmt.Printf("XML Status: %d\n", res.StatusCode) // 200
	// 响应body
	fmt.Printf("XML Body: %s", res.String())
	/*
		XML Body: {
					"args": {},
					"data": "<User><name>dj</name><age>18</age></User>",
					"files": {},
					"form": {},
					"headers": {
						"Accept-Encoding": "gzip",
						"Content-Length": "41",
						"Content-Type": "application/xml",
						"Host": "httpbin.org",
						"User-Agent": "gentleman/2.0.5",
						"X-Amzn-Trace-Id": "Root=1-624462ce-186d339f2fd3e1f774fb15e4"
					},
					"json": null,
					"origin": "121.207.135.241",
					"url": "http://httpbin.org/post"
					}
	*/
}

/*
	content-type：application/xml
*/
func GentlemanCompositionViaMultiplexer(Url string) { // "http://httpbin.org"
	// 创建客户端（实例化）
	cli := gentleman.New()

	// 定义服务器Url(必须是第一个)
	cli.Use(url.URL(Url))

	// 基于多个匹配器创建一个新的多路复用器
	mx := mux.If(mux.Method("GET"), mux.Host("httpbin.org"))

	// 在多路复用器上附加一个自定义插件，如果匹配器通过将执行该插件
	mx.Use(url.Path("/headers"))

	// 在主客户端连接多路复用器
	cli.Use(mx)

	// 指定请求发送
	res, err := cli.Request().Send()

	// 判断是否请求异常
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	// 判断是否请求失败（请求成功 res.Ok=true）
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	// 响应状态码
	fmt.Printf("Status: %d\n", res.StatusCode)
	// 响应body
	fmt.Printf("Body: %s", res.String())
}

func QueryUrlParamRequest() {
	cli := gentleman.New()
	cli.URL("https://api.thecatapi.com/")

	cli.Use(headers.Set("x-api-key", "479ce48d-db30-46a4-b1a0-91ac4c1477b8"))
	cli.Use(url.Path("/v1/:type"))
	// 遍历命令行入参 
	for _, arg := range os.Args[1:] {
		//  基于客户端创建心情求
		req := cli.Request()
		// 将参数加入到url
		req.Use(url.Param("type", arg))// /v1/breeds、/v1/votes、/v1/categories
		// 执行 请求
		res, err := req.Send()
		if err != nil {
			fmt.Printf("Request error: %s\n", err)
			return
		}
		if !res.Ok {
			fmt.Printf("Invalid server response: %d\n", res.StatusCode)
			return
		}

		fmt.Printf("Query  Status: %d\n", res.StatusCode)
		fmt.Printf("Query Body: %s\n", res.String())
	}
	/*
	命令行执行：
	go run main.go breeds votes categories
	*/
}
