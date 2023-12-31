package bithumb

import "github.com/valyala/fasthttp"

type BithumbRestConnector struct {
	RestConnector
	client *fasthttp.Client
	host   string
}

// NewBithumbRestConnector 빗썸 커넥터 생성자
func NewBithumbRestConnector() *BithumbRestConnector {
	client := new(fasthttp.Client)
	client.DisableHeaderNamesNormalizing = true
	return &BithumbRestConnector{
		client: client,
		host:   "https://api.bithumb.com",
	}
}
