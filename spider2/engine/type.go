package engine

import "net/http"

/**
自定义请求
 */
type CRequest struct {
	Req http.Request
	ParseFunc func([]byte)()
}

/**
自定义返回结果
 */
type CResponse struct {
	CRequests []CRequest

}

