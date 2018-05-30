package collector

import (
	"monitoring/core"
	"monitoring/internal"
	"github.com/valyala/fasthttp"
	"fmt"
	"os"
)

type Metric struct {
	*Request
	Client *fasthttp.Client
	Config *internal.Configuration
	System *core.System
}

func NewMetric(system *core.System, config *internal.Configuration) *Metric {
	client := &fasthttp.Client{}

	metric := &Metric{
		Request:&Request{},
		System:system,
		Config:config,
		Client:client,
	}

	return metric
}

func (r *Metric)push(c chan bool) {
	for {
		<-c
		r.Request = NewRequest(r.System, r.Config.Addr)
		r.do(r)
		fmt.Fprintln(os.Stdout, r.HTTPResponse)
		fasthttp.ReleaseRequest(r.HTTPRequest)
		fasthttp.ReleaseResponse(r.HTTPResponse)
	}
}
