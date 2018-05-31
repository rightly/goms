package collector

import (
	"monitoring/core"
	"monitoring/internal"
	"sync"
	"github.com/valyala/fasthttp"
	"fmt"
	"os"
)

type collector struct {
	*request
	System *core.System
	Client *fasthttp.Client
	Config *internal.Configuration
}

func New(system *core.System, config *internal.Configuration) *collector {
	client := &fasthttp.Client{}

	return &collector{
		System: system,
		Client: client,
		Config: config,
	}
}

func (c *collector)Start(wg *sync.WaitGroup)  {
	collectChan := make(chan bool)
	var wait sync.WaitGroup

	wait.Add(2)

	go c.push(collectChan)
	go c.System.Collect(collectChan)

	wait.Wait()

	wg.Done()
}

func (c *collector)push(ch chan bool) {
	for {
		<-ch
		c.request = newRequest(c.Config, c.System)
		c.do()
		fmt.Fprintln(os.Stdout, c.HTTPResponse)
		fasthttp.ReleaseRequest(c.HTTPRequest)
		fasthttp.ReleaseResponse(c.HTTPResponse)
	}
}

func (c *collector)do() error {
	err := c.Client.Do(c.request.HTTPRequest, c.request.HTTPResponse)
	if internal.CheckErr(err, "couldn't send request in collector") {
		return err
	}

	return nil
}