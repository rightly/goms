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

func (c *manager)init() {
	// Echo init
	fp, _ := os.OpenFile("./echo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	c.web = echo.New()

	c.web.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:fp,
	}))

	// middleware
	c.web.Use(middleware.Recover())
	c.web.Use(middleware.CORS())
	c.web.Use(middleware.RequestID())
	c.web.HTTPErrorHandler = errorHandler

	// goms admin page 핸들러 추가
	c.web.GET("/", index)

	// 각 API 리소스별 핸들러 추가
	metricHandler(c.web)

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