package httpclient

import (
	"fmt"
	"testing"
)

func TestHttpClient(t *testing.T) {
	req := Request{
		URL: "https://pokeapi.co/api/v2/ability/?offset=0&limit=1",
	}
	resp := CheckUptime(req)
	fmt.Println("Response : ", resp)
}
