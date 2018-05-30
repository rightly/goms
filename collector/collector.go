package collector

import (
	"monitoring/core"
	"monitoring/internal"
	"sync"
)

type Collector struct {
	*Metric	`json:"metric"`
}

func New(system *core.System, config *internal.Configuration) *Collector {
	metric := NewMetric(system, config)
	return &Collector{
		Metric: metric,
	}
}

func (c *Collector)Start(wg *sync.WaitGroup)  {
	collectChan := make(chan bool)
	var wait sync.WaitGroup

	wait.Add(2)

	go c.push(collectChan)
	go c.System.Collect(collectChan)

	wait.Wait()

	wg.Done()
}