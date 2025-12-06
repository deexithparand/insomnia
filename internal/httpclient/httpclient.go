package httpclient

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Status       string        `json:"status"`
	StatusCode   int           `json:"status_code"`
	ResponseTime time.Duration `json:"response_time"`
	Body         string        `json:"body"`
}

func CheckUptime(req Request) Response {

	var (
		uptimeResp             Response
		httpBufferedLineOutput string
		timeBeforeRequest      time.Time = time.Now()
	)

	httpResp, err := http.Get(req.URL)
	if err != nil {
		panic(err)
	}

	// reader to read data from the bufio reader
	bufioReader := bufio.NewReader(httpResp.Body)
	defer httpResp.Body.Close()
	for {
		httpLineOutput, err := bufioReader.ReadString('\n')

		if len(httpLineOutput) > 0 {
			httpBufferedLineOutput += httpLineOutput
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error reading from the http response body : ", err)
		}
	}

	// update values
	uptimeResp.Status = httpResp.Status
	uptimeResp.StatusCode = httpResp.StatusCode
	uptimeResp.Body = httpBufferedLineOutput
	uptimeResp.ResponseTime = time.Since(timeBeforeRequest)

	return uptimeResp
}
