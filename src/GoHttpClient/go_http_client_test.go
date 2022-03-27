package GoHttpClient

import (
	"testing"
)

func TestGoHttpClientA(t *testing.T) {
	// GoNativeHttpClientGet() // Go native
	GentlemanSampleRequest("http://httpbin.org/post") // Go gentleman
	
}


