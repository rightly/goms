package collector

import (
	"github.com/valyala/fasthttp"
	"monitoring/core"
	"encoding/json"
	"monitoring/internal"
)

type request struct {
	RequestContext *fasthttp.RequestCtx
	HTTPRequest    *fasthttp.Request
	HTTPResponse   *fasthttp.Response
	Result         string
	Body           []byte
	Data           *core.System
}

func newRequest(config *internal.Configuration, system *core.System) *request {
	protocol := "http://"
	endPoint := "/metrics"
	url := protocol + config.Addr + endPoint
	method := "PUT"
	body, err := json.Marshal(system)
	internal.CheckErr(err, "couldn't marshal core.System in Request.New")

	fastRequest := fasthttp.AcquireRequest()
	fastResponse := fasthttp.AcquireResponse()

	fastRequest.SetRequestURI(url)
	fastRequest.SetBody(body)
	fastRequest.Header.SetContentType("application/json")

	setDefaultHeader(fastRequest, config.Name, method)

	req := &request{
		HTTPRequest:fastRequest,
		HTTPResponse:fastResponse,
		Data: &core.System{},
	}

	return req
}

func setDefaultHeader(r *fasthttp.Request, xServer, method string )  {
	r.Header.SetMethod(method)
	r.Header.SetUserAgent("Goms")
	r.Header.Set("Server", "Goms-collector/0.1")
	r.Header.Set("X-Server", xServer)
}