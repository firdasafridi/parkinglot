package reqinfo

import (
	"context"
	"time"
)

type RequestInfo struct {
	StartRequest time.Time
	Host         string
	SourceIP     string
	RequestURL   string
	Method       string
	Error        error
	HTTPStatus   int
	UserAgent    string
}

var timeNow = time.Now

type ctxRequestInfo struct{}

// GetLatency is function get latency raw
func (req *RequestInfo) GetLatency() (latency float64) {
	latency = float64(time.Since(req.StartRequest).Milliseconds())
	return
}

// GetRequestInfo is function to get request info
func GetRequestInfo(ctx context.Context) RequestInfo {
	requestInfo, ok := ctx.Value(ctxRequestInfo{}).(RequestInfo)
	if !ok {
		return RequestInfo{}
	}
	return requestInfo
}
