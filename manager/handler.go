package manager

import (
	"github.com/labstack/echo"
	"net/http"
)

func setDefaultHeader(ctx *echo.Context) {
	c := *ctx
	c.Response().Header().Set("Server", "goms-manager/0.1")
	c.Response().Header().Del("Vary")
}

// Handler
func createMetric(c echo.Context) error {
	setDefaultHeader(&c)
	remote := c.Request().Header.Get("X-Server")
	res := ok(remote)

	c.Response().Header().Set("Cache-Control", "no-cache, no-store")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	return c.JSON(http.StatusCreated, res)
}