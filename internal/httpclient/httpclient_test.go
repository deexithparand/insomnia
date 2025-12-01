package httpclient

import (
	"fmt"
	"testing"
)

func TestHttpClient(t *testing.T) {
	req := Request{
		URL: "https://www.google.com",
	}
	resp := GETTargetUptime(req)
	fmt.Println(resp.StatusCode)
	// fmt.Println("he;lo")
}
