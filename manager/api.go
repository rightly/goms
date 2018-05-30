package manager

import (
	"monitoring/core"
	"github.com/valyala/fasthttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)
type Application string

type Collector struct {
	Application
	*core.System
	Request *fasthttp.Request
}

func (c *Collector)Send() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}