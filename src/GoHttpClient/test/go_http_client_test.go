package test

import (
	"testing"
	GoHttpClient "src/GoHttpClient"
)

func TestGoHttpClientA(t *testing.T) {
	// GoNativeHttpClientGet() // Go native
	GoHttpClient.GentlemanSampleRequest("http://httpbin.org/post") // Go gentleman

}

/*
	命令行执行：
	cd /goland-learning/src/GoHttpClient/test
	go test -v -run TestGoHttpClientA
	go test -v -run TestGoHttpClientA go_http_client_test.go
	go test -v go_http_client_test.go
*/