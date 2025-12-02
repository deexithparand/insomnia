package httpclient

import (
	"fmt"
	"testing"
)

func TestHttpClient(t *testing.T) {
	req := Request{
		URL: "https://pokeapi.co/api/v2",
	}
	resp := GETTargetUptime(req)
	fmt.Println(resp.StatusCode)
}
