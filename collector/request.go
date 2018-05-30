package collector

import (
	"github.com/valyala/fasthttp"
	"monitoring/core"
	"encoding/json"
	"monitoring/internal"
)

type Request struct {
	RequestContext *fasthttp.RequestCtx
	HTTPRequest    *fasthttp.Request
	HTTPResponse   *fasthttp.Response
	Result         string
	Body           []byte
	Data           *core.System
}

func NewRequest(r *core.System, manager string) *Request {
	protocol := "http://"
	endPoint := "/push-metric"
	url := protocol + manager + endPoint
	method := "PUT"
	body, err := json.Marshal(r)
	internal.CheckErr(err, "couldn't marshal core.System in Request.New")


	fastRequest := fasthttp.AcquireRequest()
	fastResponse := fasthttp.AcquireResponse()

	fastRequest.SetRequestURI(url)
	fastRequest.SetBody(body)

	fastRequest.Header.SetMethod(method)
	fastRequest.Header.SetContentType("application/json")
	fastRequest.Header.SetUserAgent("Goms")
	fastRequest.Header.Set("Server", "Goms-collector/0.1")

	req := &Request{
		HTTPRequest:fastRequest,
		HTTPResponse:fastResponse,
		Data: &core.System{},
	}

	return req
}

func (r *Request)do(m *Metric) error {
	//dial := fasthttp.Dial()

	err := m.Client.Do(r.HTTPRequest, r.HTTPResponse)
	if internal.CheckErr(err, "couldn't send request in collector") {
		return err
	}

	return nil
}