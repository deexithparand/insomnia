package internal

import "time"

// request model
type Request struct {
	Interval          string    `json:"interval"`
	Endpoint          string    `json:"url"`
	LastAccessTime    time.Time `json:"last-access-time"`
	CurrentAccessTime time.Time `json:"current-access-time"`
}

// response model
type Response struct {
	Interval          string    `json:"interval"`
	Endpoint          string    `json:"url"`
	LastAccessTime    time.Time `json:"last-access-time"`
	CurrentAccessTime time.Time `json:"current-access-time"`
	NextAccessTime    time.Time `json:"next-access-time"`
	ResponseTime      string    `json:"response-time"`
	StatusCode        string    `json:"status-code"`
}
