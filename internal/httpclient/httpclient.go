package httpclient

import (
	"net/http"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	// ResponseTime int    `json:"response_time"`
	// Body string `json:"body"`
}

func GETTargetUptime(req Request) Response {

	var uptimeResp Response

	httpResp, err := http.Get(req.URL)
	if err != nil {
		panic(err)
	}

	// update values
	uptimeResp.Status = httpResp.Status
	uptimeResp.StatusCode = httpResp.StatusCode

	// var respBody string = ""
	// bufferedReader := bufio.NewReader(httpResp.Body)
	// for {
	// 	line, err := bufferedReader.ReadString('\n')
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	respBody += line
	// }

	// uptimeResp.Body = respBody

	return uptimeResp
}
