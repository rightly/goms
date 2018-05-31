package manager

import (
	"sync"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"os"
	"github.com/labstack/echo/middleware"
	"monitoring/internal"
	"strconv"
)

type manager struct {
	db     *xorm.Engine
	web    *echo.Echo
	config *internal.Configuration
}

func (r *manager)init() {
	//Echo init
	fp, _ := os.OpenFile("./echo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	r.web = echo.New()

	r.web.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:fp,
	}))

	//middleware
	r.web.Use(middleware.Recover())
	r.web.Use(middleware.RequestID())
	r.web.Use(middleware.CORS())

	r.web.PUT("/metric", createMetric)
	//xorm init

}

func New(config *internal.Configuration) *manager {
	manager := &manager{
		db:     &xorm.Engine{},
		web:    &echo.Echo{},
		config: config,
	}

	manager.init()
	
	return manager
}

func (c *manager)Start(wg *sync.WaitGroup)  {
	port := ":" + strconv.FormatUint(uint64(c.config.Port), 10)
	// Start server
	c.web.Logger.Fatal(c.web.Start(port))
	wg.Done()
}